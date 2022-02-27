package myhandler

import (
	"net/http"
	"time"

	"github.com/time2k/letsgo-ng"

	myconfig "mypws/config"
	mymodel "mypws/model"
	mytypedef "mypws/typedef"
	"regexp"
)

//AddDevice 苹果支付回调验证
func AddDevice(commp letsgo.CommonParams) error {
	//通用参数处理，通用参数包括letsgo框架指针通过此结构体传递到model
	c := commp.HTTPContext

	reqParams := letsgo.ParamTrim(c.QueryParam("devicename"))
	devicename := reqParams[0]

	//查询设备名是否合法及被占用
	ifmatch, _ := regexp.MatchString("[a-z0-9]+", devicename)
	if !ifmatch {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "devicename invalid", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusBadRequest, ret.FormatNew())
	}

	if mymodel.DeviceNameCheckExists(commp, devicename) {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "devicename exists!", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusBadRequest, ret.FormatNew())
	}

	deviceinfo := mytypedef.PWSDeviceInfo{}

	deviceinfo.DeviceName = devicename
	deviceinfo.DeviceID = mymodel.GenDeviceID()
	deviceinfo.Password = mymodel.GenPassword()

	ret := mymodel.InsertDeviceInfo(commp, deviceinfo)

	return c.JSON(http.StatusOK, ret.FormatNew())
}

//AddData 添加数据
func AddData(commp letsgo.CommonParams) error {
	//通用参数处理，通用参数包括letsgo框架指针通过此结构体传递到model
	c := commp.HTTPContext

	u := new(mytypedef.PWSData)
	if err := c.Bind(u); err != nil {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: err.Error(), Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusBadRequest, ret.FormatNew())
	}
	if err := c.Validate(u); err != nil {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: err.Error(), Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusBadRequest, ret.FormatNew())
	}

	ret := mymodel.DeviceInfo(commp, u.ID, u.PASSWORD)
	if ret.Status != myconfig.StatusOk {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "ID,PASSWORD Invalid", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusBadRequest, ret.FormatNew())
	}
	deviceinfo := ret.Body.(mytypedef.PWSDeviceInfo)

	if u.Dateutc == "now" {
		u.Dateutc = time.Now().Format("2006-01-02 15:04:05")
	}

	//https://support.weather.com/s/article/PWS-Upload-Protocol?language=en_US

	ret2 := mymodel.InsertData(commp, deviceinfo.DeviceName, *u)

	return c.JSON(http.StatusOK, ret2.FormatNew())
}

func GetData(commp letsgo.CommonParams) error {
	//通用参数处理，通用参数包括letsgo框架指针通过此结构体传递到model
	c := commp.HTTPContext

	reqParams := letsgo.ParamTrim(c.QueryParam("devicename"), c.QueryParam("interval"))
	devicename := reqParams[0]
	interval := reqParams[1]

	//查询设备名是否合法及被占用
	ifmatch, _ := regexp.MatchString("[a-z0-9]+", devicename)
	if !ifmatch {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "devicename invalid", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusBadRequest, ret.FormatNew())
	}

	if interval == "realtime" {
		ret := mymodel.SelectRealtimeData(commp, devicename)
		return c.JSON(http.StatusOK, ret.FormatNew())
	}

	return c.JSON(http.StatusOK, nil)
}
