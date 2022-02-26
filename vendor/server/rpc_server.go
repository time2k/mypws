package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/time2k/letsgo-ng"

	//"net/http"
	myconfig "social_pay/config"
	myhandler "social_pay/handler"
)

func main() {
	log.Println(myconfig.PROJECT_NAME, "RPC Server Start")

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

	msclient := letsgo.NewConsulClient() //使用consul作为微服务服务框架
	letsgo.Default.InitMicroserviceClient(msclient)
	letsgo.Default.MicroserviceClient.RegisterService(myconfig.PROJECT_NAME+"-rpc", myconfig.RPC_SERVER_PORT)

	defer letsgo.Default.Close()

	file, err := os.OpenFile(myconfig.RPC_ACCESS_LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rpc_logger := log.New(file, "", log.LstdFlags)

	commp := letsgo.CommonParams{}
	commp.Init()

	prpc := new(myhandler.PRPC)
	prpc.Commp = commp
	prpc.Logger = rpc_logger

	server := rpc.NewServer()
	server.Register(prpc)

	ln, e := net.Listen("tcp", fmt.Sprint(":", myconfig.RPC_SERVER_PORT))
	if e != nil {
		panic(e)
	}
	defer ln.Close()

	go func() {
		for {
			conn, e := ln.Accept()
			if e != nil {
				continue
			}
			rpc_logger.Println(conn.RemoteAddr().String())
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()

	// Wait for termination signal
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptSignal

	// Terminate the server
	log.Println(myconfig.PROJECT_NAME, "RPC Server Stop")
}
