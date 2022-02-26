package myhandler

import (
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/time2k/letsgo-ng"

	myconfig "mypws/config"
	mymodel "mypws/model"
	mytypedef "mypws/typedef"
	"regexp"
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

//AddDevice 苹果支付回调验证
func AddDevice(commp letsgo.CommonParams) error {
	//通用参数处理，通用参数包括letsgo框架指针通过此结构体传递到model
	c := commp.HTTPContext

	reqParams := letsgo.ParamTrim(c.QueryParam("devicename"))
	devicename := reqParams[0]

	//查询设备名是否合法及被占用
	ifmatch, _ := regexp.MatchString("[a-z0-9]+", devicename)
	if !ifmatch {
		return c.String(http.StatusBadRequest, "devicename invalid")
	}

	if mymodel.DeviceNameCheckExists(commp, devicename) {
		return c.String(http.StatusBadRequest, "devicename exists")
	}

	deviceid := mymodel.GenDeviceID()
	passwd := mymodel.GenPassword()

	mymodel.InsertDeviceInfo(commp, devicename, deviceid, passwd)

	return c.String(http.StatusOK, "success"+" "+deviceid+" "+passwd)
}

//AddData 添加数据
func AddData(commp letsgo.CommonParams) error {
	//通用参数处理，通用参数包括letsgo框架指针通过此结构体传递到model
	c := commp.HTTPContext

	u := new(mytypedef.PWSData)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	ret := mymodel.DeviceInfo(commp, u.ID, u.PASSWORD)
	if ret.Status != myconfig.StatusOk {
		return c.String(http.StatusBadRequest, "ID PASSWORD not match")
	}
	deviceinfo := ret.Body.(mytypedef.PWSDeviceInfo)

	if u.Dateutc == "now" {
		u.Dateutc = time.Now().Format("2006-01-02 15:04:05")
	}

	//https://support.weather.com/s/article/PWS-Upload-Protocol?language=en_US

	mymodel.InsertData(commp, deviceinfo.DeviceName, *u)

	return c.String(http.StatusOK, "success")
}
