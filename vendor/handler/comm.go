package myhandler

import (
	myconfig "social_pay/config"
	mytypedef "social_pay/typedef"

	"github.com/time2k/letsgo-ng"
)

//CommRiskControllWithUser 通用风控
func CommRiskControllWithUser(commp letsgo.CommonParams, accesstoken string) letsgo.BaseReturnData {
	rpc_client, err := letsgo.Default.JSONRPCClient.Dial("ucenter")
	if err != nil {
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "RPC调用错误", Body: err.Error(), IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	defer rpc_client.Close()

	//风控验证
	var data mytypedef.URPCRiskCheckRequest
	var retrpc mytypedef.URPCRiskCheckResponse
	data = mytypedef.URPCRiskCheckRequest(commp.GetParam("did"))
	err_rpc := rpc_client.Call("URPC.RiskCheck", data, &retrpc)
	if err_rpc != nil {
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "RPC调用错误", Body: err_rpc.Error(), IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	if retrpc == -2 { //devid被封，返回2008使accesstoken失效
		return letsgo.BaseReturnData{Status: 2008, Msg: "accesstoken无效", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	var data2 mytypedef.URPCValidateRequest
	var retrpc2 mytypedef.URPCValidateResponse
	data2 = mytypedef.URPCValidateRequest(accesstoken)
	err_rpc = rpc_client.Call("URPC.Validate", data2, &retrpc2)
	if err_rpc != nil {
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "RPC调用错误", Body: err_rpc.Error(), IsDebug: commp.GetParam("debug"), DebugInfo: nil}

	}
	userid := int(retrpc2)
	if userid == 0 { //无效
		return letsgo.BaseReturnData{Status: 2008, Msg: "accesstoken无效", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	if userid == -1 { //过期
		return letsgo.BaseReturnData{Status: 2009, Msg: "accesstoken已过期", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "ok", Body: userid, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}
