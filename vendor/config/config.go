// +build !debug
// +build !release

package myconfig

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/time2k/letsgo-ng/config"
)

var (
	RUN_ENV string = "debug"

	LOGROOT = "/tmp"

	BASEAUTH = false

	//MySQL数据库相关配置
	DBconfigSet config.DBconfigStruct = config.DBconfigStruct{
		"pay": {
			"master": config.DBconfig{DBhostsip: "10.124.65.163:3306", DBusername: "shanyin", DBpassword: "iNhrW@UGJ5Fju9Xx", DBname: "social_pay", DBcharset: "utf8mb4", DBconnMaxConns: 20, DBconnMaxIdles: 2, DBconnMaxLifeTime: 1800 * time.Second},
			//"master": config.DBconfig{DBhostsip: "10.124.65.159:3306", DBusername: "root", DBpassword: "pay123456", DBname: "go_pay", DBcharset: "utf8", DBconnMaxConns: 20, DBconnMaxIdles: 2, DBconnMaxLifeTime: 1800 * time.Second},
			"slave": config.DBconfig{}, //如果为主从式，按照master格式填写从库信息
		},
	}

	//redis cluster相关
	REDIS_SERVER  []string           = []string{"10.124.132.46:6509", "10.124.132.47:6509", "10.124.132.112:6509"}
	REDIS_OPTIONS []redis.DialOption = []redis.DialOption{redis.DialConnectTimeout(5 * time.Second), redis.DialPassword(""), redis.DialReadTimeout(5 * time.Second)}

	//RPC server相关配置
	RPC_SERVICE map[string]config.RPCconfig = map[string]config.RPCconfig{
		"ucenter": {Network: "tcp", Address: "10.124.65.163:7101", MicroserviceName: "social-ucenter-rpc"},
	}
)

var (
	WXNotifyUrl      = "http://test.api.social.vaas.com/pay/wxPayNotify"
	ALINotifyUrl     = "http://test.api.social.vaas.com/pay/aliPayNotify"
	SalesCallbackUrl = "http://test.api.social.vaas.com/sales/recharge/paycallback"
)
