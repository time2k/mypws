package mymodel

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	myconfig "social_pay/config"
	mylibs "social_pay/libs"
	mytypedef "social_pay/typedef"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/time2k/letsgo-ng"

	"github.com/shopspring/decimal"
)

//CheckOutTradeNo 校验业务方订单号是否重复下单
func CheckOutTradeNo(commp letsgo.CommonParams, outTradeNo string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := new(mytypedef.PayOrder)

	//MySQL数据库样例 组合成主键
	cache_key := "social:check_out_trade_no_" + outTradeNo

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //1秒钟超时
	dbq.SetSQL("SELECT id,order_no,business_id,business_name,business_uid,product_id,product_name,product_desc,out_trade_no,total_fee,order_status,notify_status,notify_url,client_ip,server_ip,sandbox,create_time,update_time FROM `pay_order` WHERE out_trade_no = ?")
	dbq.SetSQLcondition(outTradeNo)
	dbq.SetResult(data) //传递指针类型struct
	dbq.SetDbname("pay")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]CheckOutTradeNo: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data_exists, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//CheckOrderSign 校验业务线签名
func CheckOrderSign(commp letsgo.CommonParams, orderParams *mytypedef.OrderCreateParams) letsgo.BaseReturnData {

	businessIdStr := strconv.Itoa(orderParams.BusinessId)
	businessUidStr := orderParams.BusinessUid
	productIdStr := orderParams.ProductId
	productNameStr := orderParams.ProductName
	productDescStr := orderParams.ProductDesc
	outTradeNoStr := orderParams.OutTradeNo
	totalFeeStr := strconv.FormatFloat(orderParams.TotalFee, 'f', -1, 64)
	notifyUrlStr := orderParams.NotifyUrl
	clientIpStr := orderParams.ClientIp
	openIdStr := orderParams.OpenId
	returnUrlStr := orderParams.ReturnUrl
	signTypeStr := orderParams.SignType
	signStr := orderParams.Sign

	//组装下单参数数组
	params := make(map[string]string)
	params["business_id"] = businessIdStr
	if businessUidStr != "" {
		params["business_uid"] = businessUidStr
	}
	params["product_id"] = productIdStr
	params["product_name"] = productNameStr
	if productDescStr != "" {
		params["product_desc"] = productDescStr
	}
	params["out_trade_no"] = outTradeNoStr
	params["total_fee"] = totalFeeStr
	params["notify_url"] = notifyUrlStr
	params["client_ip"] = clientIpStr
	if openIdStr != "" {
		params["open_id"] = openIdStr
	}
	if returnUrlStr != "" {
		params["return_url"] = returnUrlStr
	}
	//params["time"] = time.Now().Format("2006-01-02 15:04:05")
	params["sign_type"] = signTypeStr
	params["sign"] = signStr
	// 创建切片
	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}
	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)
	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		if len(params[k]) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params[k])
			buf.WriteString(`&`)
		}
	}

	businessInfo := GetBusinessInfo(commp, orderParams.BusinessId)
	if businessInfo.Status == myconfig.StatusNoData {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)
	if businessData.BusinessKey == "" {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线加密Key", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}

	}
	businessKey := businessData.BusinessKey
	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(businessKey)
	log.Println("PARAMS SIGN DATA:", buf.String())
	var (
		dataMd5    [16]byte
		dataSha256 []byte
		str        string
	)
	//加密方式
	switch signTypeStr {
	case "MD5":
		dataMd5 = md5.Sum(buf.Bytes())
		str = hex.EncodeToString(dataMd5[:]) //需转换成切片
	case "HMAC-SHA256":
		h := hmac.New(sha256.New, []byte(businessKey))
		h.Write(buf.Bytes())
		dataSha256 = h.Sum(nil)
		str = hex.EncodeToString(dataSha256[:])
	}
	//校验参数签名是否正确
	sign := strings.ToUpper(str)
	log.Println("PARAMS SIGN:", signStr)
	log.Println("COMPUTE SIGN:", sign)
	if sign == signStr {
		return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: true, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: false, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//InsertPayOrder 存储订单信息
func InsertPayOrder(commp letsgo.CommonParams, orderParams *mytypedef.OrderCreateParams) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	orderNo := MakeRandomNo(20)
	businessInfo := GetBusinessInfo(commp, orderParams.BusinessId)
	if businessInfo.Status == myconfig.StatusNoData {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()

	dbq.SetSQL("INSERT INTO `pay_order` SET order_no = ?,business_id = ?,business_name = ?,business_uid = ?,product_id = ?,product_name = ?,product_desc = ?,out_trade_no = ?,total_fee = ?,order_status = ?,notify_status = ?,notify_url = ?,client_ip = ?,server_ip = ?,sandbox = ?,create_time = ?")
	dbq.SetSQLcondition(orderNo)
	dbq.SetSQLcondition(orderParams.BusinessId)
	dbq.SetSQLcondition(businessData.BusinessName)
	dbq.SetSQLcondition(orderParams.BusinessUid)
	dbq.SetSQLcondition(orderParams.ProductId)
	dbq.SetSQLcondition(orderParams.ProductName)
	dbq.SetSQLcondition(orderParams.ProductDesc)
	dbq.SetSQLcondition(orderParams.OutTradeNo)
	dbq.SetSQLcondition(orderParams.TotalFee)
	dbq.SetSQLcondition(0)
	dbq.SetSQLcondition(0)
	dbq.SetSQLcondition(orderParams.NotifyUrl)
	dbq.SetSQLcondition(orderParams.ClientIp)
	dbq.SetSQLcondition(commp.GetParam("ip"))
	dbq.SetSQLcondition(myconfig.SandboxFalse)
	dbq.SetSQLcondition(time.Now().Unix())
	dbq.SetDbname("pay")
	//使用多条SQL查询
	orderId, err := letsgo.Default.DBQuery.EXEC(dbq)
	if err != nil {
		letsgo.Default.Logger.Printf("[Model]InsertPayOrder: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "插入数据失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if orderId == 0 {
		letsgo.Default.Logger.Printf("[Model]InsertPayOrder: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "插入数据失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}
	//建立缓存
	PayOrderInfo(commp, orderId)

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: orderId, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//GetBusinessInfo 获取业务方配置信息
func GetBusinessInfo(commp letsgo.CommonParams, businessId int) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := new(mytypedef.BusinessPayInfo)

	//MySQL数据库样例 组合成主键
	cache_key := "social:business_pay_info_" + strconv.Itoa(businessId)

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //1秒钟超时
	dbq.SetSQL("SELECT id,business_id,business_paychannel,business_name,business_key,alipay_app_id,alipay_product_code,alipay_notify_url,alipay_private_key,alipay_public_key,wx_app_id,wx_mch_id,wx_trade_type,wx_sign_key,wx_notify_url,wx_body,status,create_time,update_time FROM `business_pay_info` WHERE business_id = ?")
	dbq.SetSQLcondition(businessId)
	dbq.SetResult(data) //传递指针类型struct
	dbq.SetDbname("pay")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]GetBusinessInfo: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)
	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线支付信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//GetBusinessInfoList 获取业务方配置信息列表 todo:业务方区分id
func GetBusinessInfoList(commp letsgo.CommonParams) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := mytypedef.BusinessPayInfoList{}

	//MySQL数据库样例 组合成主键
	cache_key := "social:business_pay_info_list"

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //1秒钟超时
	dbq.SetSQL("SELECT id,business_id,business_paychannel,business_name,business_key,alipay_app_id,alipay_product_code,alipay_notify_url,alipay_private_key,alipay_public_key,wx_app_id,wx_mch_id,wx_trade_type,wx_sign_key,wx_notify_url,wx_body,status,create_time,update_time FROM `business_pay_info`")
	dbq.SetResult(&data.List) //传递指针类型struct
	dbq.SetDbname("pay")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectMulti(dbq)
	if err != nil {
		letsgo.Default.Logger.Panicf("[Model]GetBusinessInfoList: %s", err.Error())
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)
	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线支付信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//PayOrderInfo 支付订单信息
func PayOrderInfo(commp letsgo.CommonParams, orderId int64) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := new(mytypedef.PayOrder)

	orderIdStr := strconv.Itoa(int(orderId))
	//MySQL数据库样例 组合成主键
	cache_key := "social:pay_order_" + orderIdStr

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //10秒钟超时
	dbq.SetSQL("SELECT id,order_no,business_id,business_name,business_uid,product_id,product_name,product_desc,out_trade_no,total_fee,order_status,notify_status,notify_url,client_ip,server_ip,sandbox,create_time,update_time FROM `pay_order` WHERE id = ?")
	dbq.SetResult(data) //传递指针类型struct
	dbq.SetSQLcondition(orderId)
	dbq.SetDbname("pay")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Printf("[Model]PayOrderInfo: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "查询失败", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "无数据", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//PayOrderInfoByNo 支付订单信息(根据order_no订单号查)
func PayOrderInfoByNo(commp letsgo.CommonParams, orderNo string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	data := new(mytypedef.PayOrder)

	//MySQL数据库样例 组合成主键
	cache_key := "social:pay_order_no_" + orderNo

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()
	dbq.UseCache = true
	dbq.SetCacheKey(cache_key)
	dbq.SetCacheExpire(600) //10秒钟超时
	dbq.SetSQL("SELECT id,order_no,business_id,business_name,business_uid,product_id,product_name,product_desc,out_trade_no,total_fee,order_status,notify_status,notify_url,client_ip,server_ip,sandbox,create_time,update_time FROM `pay_order` WHERE order_no = ?")
	dbq.SetResult(data) //传递指针类型struct
	dbq.SetSQLcondition(orderNo)
	dbq.SetDbname("pay")
	//使用多条SQL查询
	data_exists, err := letsgo.Default.DBQuery.SelectOne(dbq)
	if err != nil {
		letsgo.Default.Logger.Printf("[Model]PayOrderInfo: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "查询失败", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if data_exists == false {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "无数据", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: data, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//PayOrderInfoUpdate 更新订单信息
func PayOrderInfoUpdate(commp letsgo.CommonParams, orderNo string, orderStatus int, notifyStatus int, isSandBox bool) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()
	if isSandBox == true {
		dbq.SetSQL("UPDATE `pay_order` SET order_status = ?, notify_status = ?, update_time = ?, sandbox = 1 WHERE order_no = ?")
	} else {
		dbq.SetSQL("UPDATE `pay_order` SET order_status = ?, notify_status = ?, update_time = ? WHERE order_no = ?")
	}
	dbq.SetSQLcondition(orderStatus)
	dbq.SetSQLcondition(notifyStatus)
	dbq.SetSQLcondition(time.Now().Unix())
	dbq.SetSQLcondition(orderNo)
	dbq.SetDbname("pay")
	affectrow, err := letsgo.Default.DBQuery.EXEC(dbq)

	if err != nil {
		letsgo.Default.Logger.Printf("[Model]PayOrderInfoUpdate: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "更新订单状态失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}
	if affectrow != 1 {
		letsgo.Default.Logger.Printf("[Model]PayOrderInfoUpdate: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "更新订单状态失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	//删除缓存
	letsgo.Default.Cache.Delete("pay_order_no_" + orderNo)

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//CallBackBusinessResult 通知业务方支付成功结果
func CallBackBusinessResult(commp letsgo.CommonParams, orderNo string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	payOrderInfo := PayOrderInfoByNo(commp, orderNo)
	log.Println("CALLBACK payOrderInfo", payOrderInfo)
	if payOrderInfo.Status != myconfig.StatusOk {
		letsgo.Default.Logger.Println("[Model]CallBackBusinessResult OrderNo err:", orderNo)
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: orderNo, Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}
	//下单时传递的notify_url
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	urlx := payOrderData.NotifyUrl + "?"

	urls := make(map[string]string)
	//urls["ticket"] = ticket
	urls["order_no"] = orderNo
	urls["pay_status"] = "1"
	urls["token"] = "K6fEPxNRIYhE"
	urls["pcode"] = "210110000"
	urls["version"] = "1.0"

	var buildURL url.URL
	q := buildURL.Query()
	for k, v := range urls {
		q.Add(k, v)
	}
	queryStr := q.Encode()

	urlx += queryStr

	headers := make(map[string]string)
	headers["Accept-Type"] = "application/json"

	//从sync poll中取一个model，比new一个model要快
	httpq := letsgo.NewHTTPQueryBuilder()
	httpq.SetCacheExpire(120) //120秒超时
	uniqid := httpq.SetRequest(true, "JSON", "GET", urlx, headers, nil)
	log.Println("CALLBACK HTTP UNIQID", uniqid)
	//运行CacheHTTP
	httpret, err := letsgo.Default.HTTPQuery.Run(httpq)
	log.Println("CALLBACK httpret:", httpret)
	if err != nil {
		letsgo.Default.Logger.Println("[Model]CallBackBusinessResult:", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: err.Error(), Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}
	debuginfo = append(debuginfo, httpq.DebugInfo)

	retjson := httpret[uniqid].(*simplejson.Json)
	log.Println("CALLBACK retjson:", retjson)
	if retjson.Get("code").MustInt() != 0 {
		//todo 通知失败 记录订单号 重新回调通知
		insertForder := InsertFailedOrder(commp, orderNo)
		if insertForder.Status != myconfig.StatusOk {
			letsgo.Default.Logger.Println("[Model]InsertFailedOrder Failed:", orderNo)
		}
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: retjson.Get("message").MustString(), Body: retjson.Get("code").MustInt(), IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//InsertFailedOrder 存储回调失败订单信息
func InsertFailedOrder(commp letsgo.CommonParams, orderNo string) letsgo.BaseReturnData {
	//初始化数据
	var debuginfo []letsgo.DebugInfo
	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()

	dbq.SetSQL("INSERT INTO `failed_notify_order` SET order_no = ?,type = ?,status = ?,failed_count = ?,create_time = ?,update_time = ?")
	dbq.SetSQLcondition(orderNo)
	dbq.SetSQLcondition(1)
	dbq.SetSQLcondition(1)
	dbq.SetSQLcondition(1)
	dbq.SetSQLcondition(time.Now().Unix())
	dbq.SetSQLcondition(time.Now().Unix())
	dbq.SetDbname("pay")
	//使用多条SQL查询
	orderId, err := letsgo.Default.DBQuery.EXEC(dbq)
	if err != nil {
		letsgo.Default.Logger.Printf("[Model]InsertFailedOrder: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "插入数据失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	debuginfo = append(debuginfo, dbq.DebugInfo)

	if orderId == 0 {
		letsgo.Default.Logger.Printf("[Model]InsertPayOrder: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "插入数据失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: orderId, IsDebug: commp.GetParam("debug"), DebugInfo: debuginfo}
}

//UpdateFailedOrder 更新回调失败订单信息
func UpdateFailedOrder(commp letsgo.CommonParams, orderNo string, status int, failedCount int) letsgo.BaseReturnData {

	//从sync poll中取一个model，比new一个model要快
	dbq := letsgo.NewDBQueryBuilder()

	if status == 0 && failedCount > 0 {
		dbq.SetSQL("UPDATE `failed_notify_order` SET failed_count = ?, update_time = ? WHERE order_no = ?")
		dbq.SetSQLcondition(failedCount)
	}
	if failedCount == 0 && status > 0 {
		dbq.SetSQL("UPDATE `failed_notify_order` SET `status` = ?, update_time = ? WHERE order_no = ?")
		dbq.SetSQLcondition(status)
	}

	dbq.SetSQLcondition(time.Now().Unix())
	dbq.SetSQLcondition(orderNo)
	dbq.SetDbname("pay")
	affectrow, err := letsgo.Default.DBQuery.EXEC(dbq)

	if err != nil {
		letsgo.Default.Logger.Printf("[Model]UpdateFailedOrder: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "更新订单状态失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	if affectrow != 1 {
		letsgo.Default.Logger.Printf("[Model]UpdateFailedOrder: %s", err.Error())
		return letsgo.BaseReturnData{Status: myconfig.StatusError, Msg: "更新订单状态失败", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	return letsgo.BaseReturnData{Status: myconfig.StatusOk, Msg: "OK", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//AliH5PayOrder 支付宝H5支付下单
func AliH5PayOrder(commp letsgo.CommonParams, orderId int64, businessId int, returnUrl string) letsgo.BaseReturnData {
	businessInfo := GetBusinessInfo(commp, businessId)
	if businessInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)
	payOrderInfo := PayOrderInfo(commp, orderId)
	if payOrderInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未找到支付订单信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	privateKey := businessData.AlipayPrivateKey
	appId := businessData.AlipayAppId
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := mylibs.NewAliPayClient(appId, privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl(myconfig.ALINotifyUrl)
	if returnUrl != "" {
		client.SetReturnUrl(returnUrl)
	}
	//请求参数
	body := make(mylibs.BodyMap)
	body.Set("subject", payOrderData.ProductName)
	body.Set("out_trade_no", payOrderData.OrderNo)
	body.Set("total_amount", payOrderData.TotalFee)
	//手机网站支付请求
	payUrl, err := client.AliPayTradeWapPay(body)
	if err != nil {
		log.Println("err:", err)
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "ALI_PAY_H5支付错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	res := mytypedef.OrderResponse{}
	res.ReturnType = "ALI_PAY_H5"
	res.ReturnData = payUrl
	res.OrderNo = payOrderData.OrderNo

	return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "ALI_PAY_H5下单成功", Body: res, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//AliAppPayOrder 支付宝APP支付下单
func AliAppPayOrder(commp letsgo.CommonParams, orderId int64, businessId int) letsgo.BaseReturnData {
	businessInfo := GetBusinessInfo(commp, businessId)
	if businessInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)
	payOrderInfo := PayOrderInfo(commp, orderId)
	if payOrderInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未找到支付订单信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	privateKey := businessData.AlipayPrivateKey
	appId := businessData.AlipayAppId
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := mylibs.NewAliPayClient(appId, privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl(myconfig.ALINotifyUrl)
	//请求参数
	body := make(mylibs.BodyMap)
	body.Set("subject", payOrderData.ProductName)
	body.Set("out_trade_no", payOrderData.OrderNo)
	body.Set("total_amount", payOrderData.TotalFee)
	//手机网站支付请求
	payUrl, err := client.AliPayTradeAppPay(body)
	if err != nil {
		log.Println("err:", err)
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "ALI_PAY_APP支付错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	res := mytypedef.OrderResponse{}
	res.ReturnType = "ALI_PAY_APP"
	res.ReturnData = payUrl
	res.OrderNo = payOrderData.OrderNo

	return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "ALI_PAY_APP下单成功", Body: res, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//WxH5PayOrder 微信H5支付下单
func WxH5PayOrder(commp letsgo.CommonParams, orderId int64, businessId int) letsgo.BaseReturnData {
	businessInfo := GetBusinessInfo(commp, businessId)
	if businessInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)
	payOrderInfo := PayOrderInfo(commp, orderId)
	if payOrderInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未找到支付订单信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	appId := businessData.WxAppId
	mchId := businessData.WxMchId
	apiKey := businessData.WxSignKey
	wxBody := businessData.WxBody
	//notifyUrl := payBusinessInfo.WxNotifyUrl
	client := mylibs.NewWeChatClient(appId, mchId, apiKey, true)
	//设置国家
	client.SetCountry(mylibs.China)
	//初始化参数Map
	body := make(mylibs.BodyMap)
	body.Set("nonce_str", mylibs.GetRandomString(32))
	body.Set("body", payOrderData.ProductName)
	body.Set("out_trade_no", payOrderData.OrderNo)
	body.Set("total_fee", decimal.NewFromFloat(payOrderData.TotalFee).Mul(decimal.NewFromFloat(100)).IntPart())
	body.Set("spbill_create_ip", payOrderData.ClientIp)
	body.Set("notify_url", myconfig.WXNotifyUrl)
	body.Set("trade_type", mylibs.TradeType_H5)
	body.Set("device_info", "WEB")
	body.Set("sign_type", mylibs.SignType_MD5)

	sceneInfo := make(map[string]map[string]string)
	h5Info := make(map[string]string)
	h5Info["type"] = "Wap"
	h5Info["wap_url"] = "www.shanyin.com"
	h5Info["wap_name"] = wxBody
	sceneInfo["h5_info"] = h5Info
	body.Set("scene_info", sceneInfo)

	//body.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	//sign := mylibs.GetWeChatParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	//sign, _ := mylibs.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	//body.Set("sign", sign)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		log.Println("err:", err)
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "WX_PAY_H5支付错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	//fmt.Println("wxRsp:", *wxRsp)
	//fmt.Println("wxRsp.MwebUrl:", wxRsp.MwebUrl)

	//timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	//获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := mylibs.GetMiniPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取H5支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := mylibs.GetH5PaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取小程序需要的paySign
	//paySign := mylibs.GetAppPaySign("wxdaa2ab9ef87b5497","", wxRsp.NonceStr, wxRsp.PrepayId, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	res := mytypedef.OrderResponse{}
	res.ReturnType = "WX_PAY_H5"
	res.ReturnData = wxRsp
	res.OrderNo = payOrderData.OrderNo

	return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "WX_PAY_H5下单成功", Body: res, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//WxPublicPayOrder 微信公众号支付下单
func WxPublicPayOrder(commp letsgo.CommonParams, orderId int64, businessId int, openId string) letsgo.BaseReturnData {
	businessInfo := GetBusinessInfo(commp, businessId)
	if businessInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)
	payOrderInfo := PayOrderInfo(commp, orderId)
	if payOrderInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未找到支付订单信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	appId := businessData.WxAppId
	mchId := businessData.WxMchId
	apiKey := businessData.WxSignKey
	client := mylibs.NewWeChatClient(appId, mchId, apiKey, true)
	//设置国家
	client.SetCountry(mylibs.China)
	//初始化参数Map
	body := make(mylibs.BodyMap)
	body.Set("nonce_str", mylibs.GetRandomString(32))
	body.Set("body", payOrderData.ProductName)
	body.Set("out_trade_no", payOrderData.OrderNo)
	body.Set("total_fee", decimal.NewFromFloat(payOrderData.TotalFee).Mul(decimal.NewFromFloat(100)).IntPart())
	body.Set("spbill_create_ip", payOrderData.ClientIp)
	body.Set("notify_url", myconfig.WXNotifyUrl)
	body.Set("trade_type", mylibs.TradeType_JsApi)
	body.Set("sign_type", mylibs.SignType_MD5)
	body.Set("openid", openId)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		log.Println("err:", err)
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "WX_PAY_JSAPI支付错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	//获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := mylibs.GetMiniPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取H5支付需要的paySign
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := mylibs.GetH5PaySign(appId, wxRsp.NonceStr, pac, mylibs.SignType_MD5, timeStamp, apiKey)
	//fmt.Println("paySign:", paySign)

	//获取小程序需要的paySign
	//paySign := mylibs.GetAppPaySign("wxdaa2ab9ef87b5497","", wxRsp.NonceStr, wxRsp.PrepayId, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	wxRsp.Sign = paySign
	wxRsp.Timestamp = timeStamp

	res := mytypedef.OrderResponse{}
	res.ReturnType = "WX_PAY_JSAPI"
	res.ReturnData = wxRsp
	res.OrderNo = payOrderData.OrderNo

	return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "WX_PAY_JSAPI下单成功", Body: res, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//WxAppPayOrder 微信APP支付下单
func WxAppPayOrder(commp letsgo.CommonParams, orderId int64, businessId int) letsgo.BaseReturnData {
	businessInfo := GetBusinessInfo(commp, businessId)
	if businessInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未配置业务线信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	businessData := businessInfo.Body.(*mytypedef.BusinessPayInfo)
	payOrderInfo := PayOrderInfo(commp, orderId)
	if payOrderInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未找到支付订单信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	appId := businessData.WxAppId
	mchId := businessData.WxMchId
	apiKey := businessData.WxSignKey
	client := mylibs.NewWeChatClient(appId, mchId, apiKey, true)
	//设置国家
	client.SetCountry(mylibs.China)
	//初始化参数Map
	body := make(mylibs.BodyMap)
	body.Set("nonce_str", mylibs.GetRandomString(32))
	body.Set("body", payOrderData.ProductName)
	body.Set("out_trade_no", payOrderData.OrderNo)
	body.Set("total_fee", decimal.NewFromFloat(payOrderData.TotalFee).Mul(decimal.NewFromFloat(100)).IntPart())
	body.Set("spbill_create_ip", payOrderData.ClientIp)
	body.Set("notify_url", myconfig.WXNotifyUrl)
	body.Set("trade_type", mylibs.TradeType_App)
	body.Set("sign_type", mylibs.SignType_MD5)

	fmt.Println(payOrderData)
	fmt.Println(body)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(body)
	if err != nil {
		log.Println("err:", err)
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "WX_PAY_APP支付错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	//获取小程序支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := mylibs.GetMiniPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取H5支付需要的paySign
	//pac := "prepay_id=" + wxRsp.PrepayId
	//paySign := mylibs.GetH5PaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, mylibs.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	//fmt.Println("paySign:", paySign)

	//获取小程序需要的paySign
	paySign := mylibs.GetAppPaySign(appId, mchId, wxRsp.NonceStr, wxRsp.PrepayId, mylibs.SignType_MD5, timeStamp, apiKey)
	//fmt.Println("paySign:", paySign)
	wxRsp.Sign = paySign
	wxRsp.Timestamp = timeStamp

	res := mytypedef.OrderResponse{}
	res.ReturnType = "WX_PAY_APP"
	res.ReturnData = wxRsp
	res.OrderNo = payOrderData.OrderNo

	return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "WX_PAY_APP下单成功", Body: res, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}

//IapPayOrder 苹果IAP支付下单
func IapPayOrder(commp letsgo.CommonParams, orderId int64) letsgo.BaseReturnData {
	payOrderInfo := PayOrderInfo(commp, orderId)
	if payOrderInfo.Status != myconfig.StatusOk {
		return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "未找到支付订单信息", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	}
	payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
	res := mytypedef.OrderResponse{}
	res.ReturnType = "IAP_PAY"
	res.ReturnData = "苹果支付方式下单成功，请等待异步回调通知支付结果"
	res.OrderNo = payOrderData.OrderNo

	return letsgo.BaseReturnData{Status: myconfig.StatusNoData, Msg: "苹果IAP下单成功", Body: res, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
}
