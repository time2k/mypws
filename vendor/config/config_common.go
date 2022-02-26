package myconfig

import "time"

var (
	PROJECT_NAME             = "social-pay"
	HTTP_SERVER_PORT         = 7006 //http服务端口号
	RPC_SERVER_PORT          = 7106 //rpc服务端口号
	HTTP_SERVER_READTIMEOUT  = 5 * time.Second
	HTTP_SERVER_WRITETIMEOUT = 30 * time.Second
	VERSION                  = ""

	LOGFULLPREFIX   = LOGROOT + "/" + PROJECT_NAME + "-"
	HTTP_ACCESS_LOG = LOGFULLPREFIX + "access.log"
	RPC_ACCESS_LOG  = LOGFULLPREFIX + "rpc-access.log"
	ERROR_LOG       = LOGFULLPREFIX + "error.log"
	HTTP_LOG        = LOGFULLPREFIX + "httpdata.log"
)

const (
	StatusOk            int = 0    //ok
	StatusNoData        int = 1002 //无数据
	StatusParamsNoValid int = 1003 //参数错误
	StatusError         int = 1004 //异常
	SandboxTrue         int = 1    //沙盒支付
	SandboxFalse        int = 2    //非沙盒支付
)

//延迟回调次数->时间
var CALLBACK_TIME = map[int]int{
	1: 15,
	2: 30,
	3: 60,
	4: 90,
	5: 120,
	6: 300,
	7: 600,
}
