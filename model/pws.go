package mymodel

import (
	myconfig "mypws/config"
	mylibs "mypws/libs"
	mytypedef "mypws/typedef"
	"strings"

	"github.com/time2k/letsgo-ng"
)

func GenDeviceID() string {
	return strings.ToUpper(mylibs.RandEngString(10))
}

func GenPassword() string {
	return strings.ToUpper(mylibs.RandEngString(10))
}

//DeviceInfo 查询设备信息
func DeviceInfo(commp letsgo.CommonParams, deviceid, password string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := mytypedef.PWSDeviceInfo{}

	//MySQL数据库样例 组合成主键
	cache_key := "mypws:deviceinfo:" + deviceid + "_" + password

	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //1秒钟超时
	dbq.SetSQL("SELECT devicename,deviceid,password FROM `pws_devices` WHERE deviceid = ? AND password = ?")
	dbq.SetSQLcondition(deviceid)
	dbq.SetSQLcondition(password)
	dbq.SetResult(&data) //传递指针类型struct
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]DeviceInfo: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "Nodata", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//DeviceNameCheckExist 查询设备名是否存在
func DeviceNameCheckExists(commp letsgo.CommonParams, devicename string) bool {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := mytypedef.PWSDeviceInfo{}

	//MySQL数据库样例 组合成主键
	cache_key := "mypws:devicename:" + devicename

	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //1秒钟超时
	dbq.SetSQL("SELECT devicename,deviceid,password FROM `pws_devices` WHERE devicename = ?")
	dbq.SetSQLcondition(devicename)
	dbq.SetResult(&data) //传递指针类型struct
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]DeviceNameCheckExists: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	return data_exists
}

//InsertDeviceInfo 插入设备信息
func InsertDeviceInfo(commp letsgo.CommonParams, devicename, deviceid, password string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	dbq := letsgo.NewDBQueryBuilder()
	dbq.SetSQL("INSERT INTO `pws_devices` (`devicename`,`deviceid`,`password`) VALUES(?,?,?)")
	dbq.SetSQLcondition(devicename)
	dbq.SetSQLcondition(deviceid)
	dbq.SetSQLcondition(password)
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	_, err := letsgo.Default.DBQuery.EXEC(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]InsertDeviceInfo: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//InsertData 插入数据
func InsertData(commp letsgo.CommonParams, devicename string, data mytypedef.PWSData) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	dbq := letsgo.NewDBQueryBuilder()
	dbq.SetSQL("INSERT INTO `pws_data` (`devicename`,`dateutc`,`createdatelocal`,`winddir`,`windspeedmph`,`windgustmph`,`windgustdir`,`windspdmph_avg2m`,`windgustmph_10m`,`windgustdir_10m`,`humidity`,`dewptf`,`tempf`,`rainin`,`dailyrainin`,`baromin`,`UV`,`solarradiation`,`indoortempf`,`indoorhumidity`,`softwaretype`) VALUES(?,?,NOW(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	dbq.SetSQLcondition(devicename)
	dbq.SetSQLcondition(data.Dateutc)
	dbq.SetSQLcondition(data.Winddir)
	dbq.SetSQLcondition(data.Windspeedmph)
	dbq.SetSQLcondition(data.Windgustmph)
	dbq.SetSQLcondition(data.Windgustdir)
	dbq.SetSQLcondition(data.Windspdmph_avg2m)
	dbq.SetSQLcondition(data.Windgustmph_10m)
	dbq.SetSQLcondition(data.Windgustdir_10m)
	dbq.SetSQLcondition(data.Humidity)
	dbq.SetSQLcondition(data.Dewptf)
	dbq.SetSQLcondition(data.Tempf)
	dbq.SetSQLcondition(data.Rainin)
	dbq.SetSQLcondition(data.Dailyrainin)
	dbq.SetSQLcondition(data.Baromin)
	dbq.SetSQLcondition(data.UV)
	dbq.SetSQLcondition(data.Solarradiation)
	dbq.SetSQLcondition(data.Indoortempf)
	dbq.SetSQLcondition(data.Indoorhumidity)
	dbq.SetSQLcondition(data.Softwaretype)
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	_, err := letsgo.Default.DBQuery.EXEC(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]InsertData: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}
