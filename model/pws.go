package mymodel

import (
	myconfig "mypws/config"
	mylibs "mypws/libs"
	mytypedef "mypws/typedef"
	"strings"
	"time"

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
func InsertDeviceInfo(commp letsgo.CommonParams, deviceinfo mytypedef.PWSDeviceInfo) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	dbq := letsgo.NewDBQueryBuilder()
	dbq.SetSQL("INSERT INTO `pws_devices` (`devicename`,`deviceid`,`password`) VALUES(?,?,?)")
	dbq.SetSQLcondition(deviceinfo.DeviceName)
	dbq.SetSQLcondition(deviceinfo.DeviceID)
	dbq.SetSQLcondition(deviceinfo.Password)
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	_, err := letsgo.Default.DBQuery.EXEC(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]InsertDeviceInfo: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: deviceinfo, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//InsertData 插入数据
func InsertData(commp letsgo.CommonParams, devicename string, data mytypedef.PWSData) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	dbq := letsgo.NewDBQueryBuilder()
	dbq.SetSQL("INSERT INTO `pws_data` (`devicename`,`dateutc`,`createdatelocal`,`winddir`,`windspeedmph`,`windgustmph`,`windgustdir`,`windspdmph_avg2m`,`windgustmph_10m`,`windgustdir_10m`,`humidity`,`dewptf`,`tempf`,`rainin`,`dailyrainin`,`baromin`,`UV`,`solarradiation`,`indoortempf`,`indoorhumidity`,`softwaretype`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	dbq.SetSQLcondition(devicename)
	dbq.SetSQLcondition(data.Dateutc)
	dbq.SetSQLcondition(data.CreateDatetime)
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

//SelectRealtimeData 查询实时数据
func SelectRealtimeData(commp letsgo.CommonParams, devicename string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := mytypedef.PWSOutputData{}

	//MySQL数据库样例 组合成主键
	cache_key := "mypws:realtimedata:" + devicename

	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(10) //10秒钟超时
	dbq.SetSQL("SELECT `dateutc`,`createdatelocal`,`winddir`,`windspeedmph`,`windgustmph`,`windgustdir`,`windspdmph_avg2m`,`windgustmph_10m`,`windgustdir_10m`,`humidity`,`dewptf`,`tempf`,`rainin`,`dailyrainin`,`baromin`,`UV`,`solarradiation`,`indoortempf`,`indoorhumidity`,`softwaretype` FROM `pws_data` WHERE devicename = ? ORDER BY ID DESC LIMIT 1")
	dbq.SetSQLcondition(devicename)
	dbq.SetResult(&data) //传递指针类型struct
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]SelectRealtimeData: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "Nodata", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//funcSelectHistoryData 查询历史数据
func SelectHistoryData(commp letsgo.CommonParams, devicename string, interval string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := mytypedef.PWSOutputDataList{}

	//MySQL数据库样例 组合成主键
	cache_key := "mypws:historydata:" + devicename + ":" + interval

	time_s := ""
	time_e := ""
	t := time.Now()
	switch interval {
	case "daily":
		time_s = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
		time_e = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05")
	case "weekly":
		ts := GetMondayOfCurrentWeek(t)
		te := GetSundayOfCurrentWeek(t)
		time_s = time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
		time_e = time.Date(te.Year(), te.Month(), te.Day(), 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05")
	case "monthly":
		ts := GetFirstDayOfMonth(t)
		te := GetLastDayOfMonth(t)
		time_s = time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
		time_e = time.Date(te.Year(), te.Month(), te.Day(), 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05")
	}

	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(1800) //1800秒钟超时
	dbq.SetSQL("SELECT CONVERT_TZ(`dateutc`,'+00:00','+08:00'),`createdatelocal`,`winddir`,`windspeedmph`,`windgustmph`,`windgustdir`,`windspdmph_avg2m`,`windgustmph_10m`,`windgustdir_10m`,`humidity`,`dewptf`,`tempf`,`rainin`,`dailyrainin`,`baromin`,`UV`,`solarradiation`,`indoortempf`,`indoorhumidity`,`softwaretype` FROM `pws_data` WHERE devicename = ? AND CONVERT_TZ(`dateutc`,'+00:00','+08:00') >= ? AND CONVERT_TZ(`dateutc`,'+00:00','+08:00') <= ? ORDER BY ID ASC")
	dbq.SetSQLcondition(devicename)
	dbq.SetSQLcondition(time_s)
	dbq.SetSQLcondition(time_e)
	dbq.SetResult(&data.List) //传递指针类型struct
	dbq.SetDbname("mypws")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectMulti(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]SelectHistoryData: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "Nodata", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

func GetFirstDayOfMonth(t time.Time) time.Time {
	// 获取指定日期所属月份的第一天0点时间
	d := t.AddDate(0, 0, -t.Day()+1)
	return d
}

func GetLastDayOfMonth(t time.Time) time.Time {
	// 获取指定日期所属月份的最后一天0点时间
	return GetFirstDayOfMonth(t).AddDate(0, 1, -1)
}

func GetMondayOfCurrentWeek(t time.Time) time.Time {
	// 获取当前周的周一
	var offset int
	if t.Weekday() == time.Sunday {
		offset = 7
	} else {
		offset = int(t.Weekday())
	}
	return t.AddDate(0, 0, -offset+1)
}

func GetSundayOfCurrentWeek(t time.Time) time.Time {
	// 获取当前周的周日
	var offset int
	if t.Weekday() == time.Sunday {
		offset = 7
	} else {
		offset = int(t.Weekday())
	}
	return t.AddDate(0, 0, 7-offset)
}
