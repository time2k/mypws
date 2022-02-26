package myhandler

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	myconfig "social_pay/config"
	mylibs "social_pay/libs"
	mymodel "social_pay/model"
	mytypedef "social_pay/typedef"
	"strconv"

	"github.com/time2k/letsgo-ng"
)

//CreateOrder 支付下单接口
func CreateOrder(commp letsgo.CommonParams) error {
	//初始化数据
	var debuginfo []letsgo.DebugInfo

	c := commp.HTTPContext

	reqParams := letsgo.ParamTrim(c.QueryParam("business_id"),
		c.QueryParam("business_uid"),
		c.QueryParam("product_id"),
		c.QueryParam("product_name"),
		c.QueryParam("product_desc"),
		c.QueryParam("out_trade_no"),
		c.QueryParam("total_fee"),
		c.QueryParam("notify_url"),
		c.QueryParam("client_ip"),
		c.QueryParam("open_id"),
		c.QueryParam("return_url"),
		c.QueryParam("sign_type"),
		c.QueryParam("sign"),
	)
	businessId, err := strconv.Atoi(reqParams[0])
	if err != nil {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "business_id参数类型错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	totalFee, err := strconv.ParseFloat(reqParams[6], 64)
	if err != nil {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "total_fee参数类型错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	//初始化参数结构
	orderParams := mytypedef.OrderCreateParams{BusinessId: businessId, BusinessUid: reqParams[1], ProductId: reqParams[2], ProductName: reqParams[3],
		ProductDesc: reqParams[4], OutTradeNo: reqParams[5], TotalFee: totalFee, NotifyUrl: reqParams[7], ClientIp: reqParams[8],
		OpenId: reqParams[9], ReturnUrl: reqParams[10], SignType: reqParams[11], Sign: reqParams[12]}
	//非空参数校验
	if orderParams.BusinessId == 0 {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "business_id参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.ProductId == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "product_id参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.ProductName == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "product_name参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.OutTradeNo == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "out_trade_no参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.TotalFee == 0 {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "total_fee参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.NotifyUrl == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "notify_url参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.ClientIp == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "client_ip参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.SignType == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "sign_type参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if orderParams.Sign == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "sign参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	//校验支付方式
	business_id := orderParams.BusinessId
	ret := mymodel.GetBusinessInfoList(letsgo.CommonParams{})
	if ret.Status != myconfig.StatusOk {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "业务配置不存在", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	businesslist := ret.Body.(mytypedef.BusinessPayInfoList).List
	business_info := mytypedef.BusinessPayInfo{}
	valid := false
	for _, eachpayinfo := range businesslist {
		if eachpayinfo.BusinessId == int64(business_id) {
			valid = true
			business_info = eachpayinfo
		}
	}
	if valid == false {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "未查询到对应business_id", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}

	//公众号&小程序支付必需open_id
	if business_info.BusinessPayChannel == "WXPAY_PUB" {
		if orderParams.OpenId == "" {
			ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "open_id参数不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
			return c.JSON(http.StatusOK, ret.FormatNew())
		}
	}

	//校验业务方订单号是否重复下单
	checkOutNo := mymodel.CheckOutTradeNo(commp, orderParams.OutTradeNo)
	if checkOutNo.Status == myconfig.StatusOk && checkOutNo.Body == true {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "out_trade_no订单号重复下单", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	debuginfo = append(debuginfo, checkOutNo.DebugInfo...)
	//校验业务线签名是否正确
	//checkSign := mymodel.CheckOrderSign(commp,&orderParams)
	//if checkSign.Status == myconfig.StatusOk && checkSign.Body == false{
	//	ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "sign加密签名错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
	//	return c.JSON(http.StatusOK, ret.FormatNew())
	//}
	//debuginfo = append(debuginfo, checkSign.DebugInfo...)
	//生成支付订单号 存储订单信息
	insertRes := mymodel.InsertPayOrder(commp, &orderParams)
	if insertRes.Status != myconfig.StatusOk {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: insertRes.Msg, Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	debuginfo = append(debuginfo, insertRes.DebugInfo...)
	orderId := insertRes.Body.(int64)

	switch business_info.BusinessPayChannel {
	case "ALIPAY_H5":
		resData := mymodel.AliH5PayOrder(commp, orderId, orderParams.BusinessId, orderParams.ReturnUrl)
		return c.JSON(http.StatusOK, resData.FormatNew())
	case "WXPAY_H5":
		resData := mymodel.WxH5PayOrder(commp, orderId, orderParams.BusinessId)
		return c.JSON(http.StatusOK, resData.FormatNew())
	case "WXPAY_PUB":
		resData := mymodel.WxPublicPayOrder(commp, orderId, orderParams.BusinessId, orderParams.OpenId)
		return c.JSON(http.StatusOK, resData.FormatNew())
	case "ALPAY_APP":
		resData := mymodel.AliAppPayOrder(commp, orderId, orderParams.BusinessId)
		return c.JSON(http.StatusOK, resData.FormatNew())
	case "WXPAY_APP":
		resData := mymodel.WxAppPayOrder(commp, orderId, orderParams.BusinessId)
		return c.JSON(http.StatusOK, resData.FormatNew())
	case "APPLEPAY":
		resData := mymodel.IapPayOrder(commp, orderId)
		return c.JSON(http.StatusOK, resData.FormatNew())
	}

	return c.JSON(http.StatusOK, nil)
}

//WxPayNotify 微信支付成功回调接口
func WxPayNotify(commp letsgo.CommonParams) error {
	//初始化数据
	c := commp.HTTPContext

	bodyData, _ := ioutil.ReadAll(c.Request().Body)
	var weChatNotifyRequest *mylibs.WeChatNotifyRequest
	err := xml.Unmarshal(bodyData, &weChatNotifyRequest)
	if err != nil {
		log.Println("WX_PAY NOTIFY XMLUnmarshal ERROR", err)
	}
	log.Println(weChatNotifyRequest)
	if err != nil {
		log.Println("WX_PAY NOTIFY REQUEST ERROR :", err)
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "微信回调请求错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	if weChatNotifyRequest.ResultCode == "SUCCESS" {
		//回调成功查询原支付订单信息
		payOrderInfo := mymodel.PayOrderInfoByNo(commp, weChatNotifyRequest.OutTradeNo)
		if payOrderInfo.Status == myconfig.StatusOk {
			payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
			if int(payOrderData.TotalFee*100) == weChatNotifyRequest.TotalFee {
				updateOrder := mymodel.PayOrderInfoUpdate(commp, weChatNotifyRequest.OutTradeNo, 1, 1, false)
				if updateOrder.Status == myconfig.StatusOk {
					//todo 异步通知第三方支付成功结果
					go mymodel.CallBackBusinessResult(commp, weChatNotifyRequest.OutTradeNo)
					rsp := new(mylibs.WeChatNotifyResponse)
					rsp.ReturnCode = mylibs.SUCCESS
					rsp.ReturnMsg = "OK"
					return c.String(http.StatusOK, rsp.ToXmlString())
				} else {
					log.Println("WX_PAY NOTIFY ORDER UPDATE FAILED", weChatNotifyRequest.OutTradeNo)
				}
			} else {
				log.Println("WX_PAY NOTIFY ORDER TOTAL_FEE NOT MATCH", weChatNotifyRequest.OutTradeNo)
			}
		} else {
			log.Println("WX_PAY NOTIFY ORDER NOT FOUND", weChatNotifyRequest.OutTradeNo)
		}
	}

	return c.String(http.StatusOK, "")
}

//AliPayNotify 支付宝支付成功回调接口
func AliPayNotify(commp letsgo.CommonParams) error {
	//初始化数据
	c := commp.HTTPContext

	var aliPayNotifyRequest *mylibs.AliPayNotifyRequest

	aliPayNotifyRequest, err := mylibs.ParseAliPayNotifyResult(c.Request())
	//fmt.Println("ALIPAY NOTIFY FORM",c.QueryParam("out_trade_no"))
	log.Println("ALIPAY NOTIFY REQUEST", aliPayNotifyRequest)
	status := "success"
	if err != nil {
		log.Println("ALI_PAY NOTIFY REQUEST ERROR :", err)
		//ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "支付宝回调请求错误", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		status = "failed"
		return c.JSON(http.StatusOK, status)
	}
	if aliPayNotifyRequest.TradeStatus == "TRADE_SUCCESS" || aliPayNotifyRequest.TradeStatus == "TRADE_FINISHED" {
		//回调成功查询原支付订单信息
		payOrderInfo := mymodel.PayOrderInfoByNo(commp, aliPayNotifyRequest.OutTradeNo)
		if payOrderInfo.Status == myconfig.StatusOk {
			payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
			totalAmount, _ := strconv.ParseFloat(aliPayNotifyRequest.TotalAmount, 64)
			if payOrderData.TotalFee == totalAmount {
				updateOrder := mymodel.PayOrderInfoUpdate(commp, aliPayNotifyRequest.OutTradeNo, 1, 1, false)
				if updateOrder.Status == myconfig.StatusOk {
					//todo 异步通知第三方支付成功结果
					go mymodel.CallBackBusinessResult(commp, aliPayNotifyRequest.OutTradeNo)
					//rsp := new(mylibs.WeChatNotifyResponse)
					//rsp.ReturnCode = mylibs.SUCCESS
					//rsp.ReturnMsg = "OK"
					//return c.String(http.StatusOK, rsp.ToXmlString())
					return c.String(http.StatusOK, status)
				} else {
					log.Println("WX_PAY NOTIFY ORDER UPDATE FAILED", aliPayNotifyRequest.OutTradeNo)
					status = "failed"
				}
			} else {
				log.Println("WX_PAY NOTIFY ORDER TOTAL_FEE NOT MATCH", aliPayNotifyRequest.OutTradeNo)
				status = "failed"
			}
		} else {
			log.Println("WX_PAY NOTIFY ORDER NOT FOUND", aliPayNotifyRequest.OutTradeNo)
			status = "failed"
		}
	}

	return c.String(http.StatusOK, status)
}

//IapPayVerify 苹果支付回调验证
func IapPayVerify(commp letsgo.CommonParams) error {
	//通用参数处理，通用参数包括letsgo框架指针通过此结构体传递到model
	c := commp.HTTPContext

	receiptData := c.QueryParam("receipt_data")                       //验证数据
	useSandbox, err := strconv.ParseBool(c.QueryParam("use_sandbox")) //是否沙盒环境
	if err != nil {
		log.Println("IAP_PAY NOTIFY USE_SANDBOX PARAM ERROR:", err)
	}
	//验证苹果支付结果
	receipt, err := mylibs.VerifyReceipt(receiptData, useSandbox)
	if err != nil {
		log.Println("IAP_PAY NOTIFY VERIFY ERROR:", err)
	}
	log.Println("IAP_PAY VERIFY RECEIPT DATA:", receipt)
	//苹果支付凭证验证成功 修改订单信息
	orderNo := c.QueryParam("order_no")
	if orderNo == "" {
		ret := letsgo.BaseReturnData{Status: myconfig.StatusParamsNoValid, Msg: "订单号不能为空", Body: nil, IsDebug: commp.GetParam("debug"), DebugInfo: nil}
		return c.JSON(http.StatusOK, ret.FormatNew())
	}
	payOrderInfo := mymodel.PayOrderInfoByNo(commp, orderNo)
	if payOrderInfo.Status == myconfig.StatusOk {
		payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
		//todo 此处暂验证苹果充值productId = 下单时的productId
		if receipt.InApp[0].ProductId == payOrderData.ProductId {
			updateOrder := mymodel.PayOrderInfoUpdate(commp, orderNo, 1, 1, useSandbox)
			if updateOrder.Status == myconfig.StatusOk {
				//todo 异步通知第三方支付成功结果
				go mymodel.CallBackBusinessResult(commp, orderNo)
				res := mytypedef.OrderResponse{}
				res.ReturnType = "IAP_PAY"
				res.ReturnData = "receipt"
				res.OrderNo = payOrderData.OrderNo
				return c.JSON(http.StatusOK, res)
			} else {
				log.Println("IAP NOTIFY ORDER UPDATE FAILED", orderNo)
			}
		}
	}

	return c.String(http.StatusOK, "")
}
