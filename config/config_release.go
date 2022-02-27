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
		"mypws": {
			"master": config.DBconfig{DBhostsip: "127.0.0.1:3306", DBusername: "mypws", DBpassword: "Mypws123!@#", DBname: "mypws", DBcharset: "utf8mb4", DBconnMaxConns: 20, DBconnMaxIdles: 2, DBconnMaxLifeTime: 1800 * time.Second},
			"slave":  config.DBconfig{}, //如果为主从式，按照master格式填写从库信息
		},
	}

	//redis cluster相关
	REDIS_SERVER  []string           = []string{"127.0.0.1:6379"}
	REDIS_OPTIONS []redis.DialOption = []redis.DialOption{redis.DialConnectTimeout(5 * time.Second), redis.DialPassword(""), redis.DialReadTimeout(5 * time.Second)}
)
