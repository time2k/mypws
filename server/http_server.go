package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/time2k/letsgo-ng"

	//"net/http"
	myconfig "mypws/config"
	myhandler "mypws/handler"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

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

	ec.Validator = &CustomValidator{validator: validator.New()}

	//跨域
	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccessControlRequestMethod, echo.HeaderAccessControlAllowHeaders},
	}))

	//初始化letsgo框架及组件
	letsgo.NewLetsgo()
	letsgo.Default.Init()
	letsgo.Default.InitRedis(1, myconfig.REDIS_SERVER, myconfig.REDIS_OPTIONS)
	letsgo.Default.InitDBQuery(myconfig.DBconfigSet)
	letsgo.Default.InitHTTPQuery(myconfig.HTTP_LOG)
	letsgo.Default.InitLog(myconfig.ERROR_LOG)
	letsgo.Default.InitSchedule()
	letsgo.Default.InitMemConfig()
	letsgo.Default.InitCacheLock()

	defer letsgo.Default.Close()

	r := ec.Group("")

	//使用gzip压缩结果
	// r.Use(middleware.Gzip())

	//路由入口
	r.GET("/device/add", letsgo.Handler(myhandler.AddDevice)) //添加设备
	r.GET("/data/add", letsgo.Handler(myhandler.AddData))     //添加数据

	//Start server
	ec.Server.Addr = fmt.Sprint(":", myconfig.HTTP_SERVER_PORT)
	ec.Server.ReadTimeout = myconfig.HTTP_SERVER_READTIMEOUT
	ec.Server.WriteTimeout = myconfig.HTTP_SERVER_WRITETIMEOUT
	//ec.Server.Handler = http.TimeoutHandler(ec.Server.Handler, 30 * time.Second, "{\"header\":{\"status\":2},\"body\":{}}")
	gracehttp.Serve(ec.Server) //block
	log.Println(myconfig.PROJECT_NAME, "HTTP Server Stop")
}
