package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	myconfig "social_pay/config"
	mymodel "social_pay/model"
	mytypedef "social_pay/typedef"
	"syscall"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/time2k/letsgo-ng"
)

func main() {
	log.Println(myconfig.PROJECT_NAME, "recallback Server Start")

	defer letsgo.PanicFunc()

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

	go func() {
		for {
			data := new(mytypedef.FailedNotifyOrderList)

			//从sync poll中取一个model，比new一个model要快
			dbq := letsgo.NewDBQueryBuilder()
			dbq.UseCache = true
			dbq.SetCacheKey("failed_notify_orders")
			dbq.SetCacheExpire(10) //1秒钟超时
			dbq.SetSQL("SELECT id,order_no,type,`status`,failed_count,create_time,update_time FROM `failed_notify_order` WHERE status = 1")
			dbq.SetResult(&data.List) //传递指针类型struct
			dbq.SetDbname("pay")
			//使用多条SQL查询
			_, err := letsgo.Default.DBQuery.SelectMulti(dbq)
			if err != nil {
				log.Printf("recallback_server, DB query: %s", err.Error())
				return
			}

			for _, v := range data.List {
				delayTime := int64(myconfig.CALLBACK_TIME[v.FailedCount])
				if time.Now().Unix() >= v.UpdateTime+delayTime && v.FailedCount <= len(myconfig.CALLBACK_TIME)+1 {
					log.Println("find failed notify order:", v.OrderNo, "create_time:", time.Unix(v.CreateTime, 0).Local().Format("2006-01-02 15:04:05"))
					payOrderInfo := mymodel.PayOrderInfoByNo(letsgo.CommonParams{}, v.OrderNo)
					if payOrderInfo.Status != myconfig.StatusOk {
						log.Printf("recallback_server, OrderInfo has no notify_url: %s", err)
						continue
					}
					//下单时传递的notify_url
					payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
					urlx := payOrderData.NotifyUrl + "?"

					urls := make(map[string]string)
					urls["order_no"] = v.OrderNo
					urls["pay_status"] = "1"
					urls["token"] = "K6fEPxNRIYhE"
					urls["pcode"] = "210110000"
					urls["version"] = "1.0"

					var buildURL url.URL
					q := buildURL.Query()
					for k, v := range urls {
						q.Add(k, v)
					}
					queryStr := q.Encode()

					urlx += queryStr

					headers := make(map[string]string)
					headers["Accept-Type"] = "application/json"

					//从sync poll中取一个model，比new一个model要快
					httpq := letsgo.NewHTTPQueryBuilder()
					httpq.SetCacheExpire(120) //120秒超时
					uniqid := httpq.SetRequest(true, "JSON", "GET", urlx, headers, nil)

					//运行CacheHTTP
					httpret, err := letsgo.Default.HTTPQuery.Run(httpq)
					if err != nil {
						log.Printf("recallback_server, CacheHTTPError: %s", err)
						continue
					}

					retjson := httpret[uniqid].(*simplejson.Json)
					if retjson.Get("code").MustInt() != 0 {
						//todo 通知失败 更新失败次数 重新回调通知
						mymodel.UpdateFailedOrder(letsgo.CommonParams{}, v.OrderNo, 0, v.FailedCount+1)
					} else {
						//todo 通知成功 更新失败订单状态为已成功
						mymodel.UpdateFailedOrder(letsgo.CommonParams{}, v.OrderNo, 2, 0)
					}

					log.Printf("done!")
				}
			}

			time.Sleep(time.Second * 5)
		}
	}()

	// Wait for termination signal
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptSignal

	// Terminate the server
	log.Println(myconfig.PROJECT_NAME, "recallback Server Stop")
}
