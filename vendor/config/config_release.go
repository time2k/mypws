// +build release

package myconfig

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/time2k/letsgo-ng/config"
)

var (
	RUN_ENV string = "release"

	LOGROOT = "/data/logs"

	BASEAUTH = false

	//MySQL数据库相关配置
	DBconfigSet config.DBconfigStruct = config.DBconfigStruct{
		"pay": {
			"master": config.DBconfig{DBhostsip: "192.168.0.12:3306", DBusername: "social_api", DBpassword: "MO0ihdxI3uES", DBname: "social_pay", DBcharset: "utf8mb4", DBconnMaxConns: 20, DBconnMaxIdles: 2, DBconnMaxLifeTime: 1800 * time.Second},
			"slave":  config.DBconfig{DBhostsip: "192.168.0.10:3306", DBusername: "social_api", DBpassword: "MO0ihdxI3uES", DBname: "social_pay", DBcharset: "utf8mb4", DBconnMaxConns: 20, DBconnMaxIdles: 2, DBconnMaxLifeTime: 1800 * time.Second}, //如果为主从式，按照master格式填写从库信息
		},
	}

	//redis cluster相关
	REDIS_SERVER  []string           = []string{"192.168.0.7:6379"}
	REDIS_OPTIONS []redis.DialOption = []redis.DialOption{redis.DialConnectTimeout(5 * time.Second), redis.DialPassword("5UAel1RJGkVj"), redis.DialReadTimeout(30 * time.Second)}

	//RPC server相关配置
	RPC_SERVICE map[string]config.RPCconfig = map[string]config.RPCconfig{
		"ucenter": {Network: "tcp", Address: "social-ucenter-rpc.service.consul:7101", MicroserviceName: "social-ucenter-rpc"},
	}
)

var (
	WXNotifyUrl      = "https://api.social.ehaonan.com/pay/wxPayNotify"
	ALINotifyUrl     = "https://api.social.ehaonan.com/pay/aliPayNotify"
	SalesCallbackUrl = "https://api.social.ehaonan.com/sales/recharge/paycallback"
)
