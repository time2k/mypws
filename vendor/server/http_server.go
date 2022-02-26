package main

import (
	"fmt"
	"log"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/time2k/letsgo-ng"

	//"net/http"
	myconfig "social_pay/config"
	myhandler "social_pay/handler"
	mymiddleware "social_pay/middleware"
)

func main() {
	log.Println(myconfig.PROJECT_NAME, "HTTP Server Start")

	defer letsgo.PanicFunc()

	//echo框架初始化及路由设置
	ec := echo.New()
	//使用logger中间件记录请求
	logfile, err := os.OpenFile(myconfig.HTTP_ACCESS_LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("Can't open the log file %s", myconfig.HTTP_ACCESS_LOG)
	}
	defer logfile.Close()

	ec.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: logfile}))

	//使用recover中间件
	ec.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisablePrintStack: false}))

	//跨域
	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"https://mp.le.com", "https://lebz.le.com"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccessControlRequestMethod, echo.HeaderAccessControlAllowHeaders},
	}))

	//初始化letsgo框架及组件
	letsgo.NewLetsgo()
	letsgo.Default.Init()
	letsgo.Default.InitRedis(myconfig.REDIS_SERVER, myconfig.REDIS_OPTIONS)
	letsgo.Default.InitDBQuery(myconfig.DBconfigSet)
	letsgo.Default.InitHTTPQuery(myconfig.HTTP_LOG)
	letsgo.Default.InitLog(myconfig.ERROR_LOG)
	letsgo.Default.InitSchedule()
	letsgo.Default.InitJSONRPC(myconfig.RPC_SERVICE)
	letsgo.Default.InitMemConfig()
	letsgo.Default.InitCacheLock()

	msclient := letsgo.NewConsulClient() //使用consul作为微服务服务框架
	letsgo.Default.InitMicroserviceClient(msclient)
	letsgo.Default.MicroserviceClient.RegisterService(myconfig.PROJECT_NAME+"-http", myconfig.HTTP_SERVER_PORT)

	defer letsgo.Default.Close()

	r := ec.Group("")

	//使用gzip压缩结果
	// r.Use(middleware.Gzip())

	//使用baseauth验证请求
	if myconfig.BASEAUTH {
		ba := mymiddleware.NewBaseAuth()
		r.Use(ba.BaseAuth)
	}

	//路由入口
	r.GET("/pay/createorder", letsgo.Handler(myhandler.CreateOrder))   //支付下单
	r.POST("/wxPayNotify", letsgo.Handler(myhandler.WxPayNotify))      //微信支付成功回调接口
	r.POST("/aliPayNotify", letsgo.Handler(myhandler.AliPayNotify))    //支付宝支付成功回调接口
	r.GET("/pay/iapPayVerify", letsgo.Handler(myhandler.IapPayVerify)) //苹果支付回调验证

	//Start server
	ec.Server.Addr = fmt.Sprint(":", myconfig.HTTP_SERVER_PORT)
	ec.Server.ReadTimeout = myconfig.HTTP_SERVER_READTIMEOUT
	ec.Server.WriteTimeout = myconfig.HTTP_SERVER_WRITETIMEOUT
	//ec.Server.Handler = http.TimeoutHandler(ec.Server.Handler, 30 * time.Second, "{\"header\":{\"status\":2},\"body\":{}}")
	gracehttp.Serve(ec.Server) //block
	log.Println(myconfig.PROJECT_NAME, "HTTP Server Stop")
}
