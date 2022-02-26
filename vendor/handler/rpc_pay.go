package myhandler

import (
	"fmt"
	"log"
	myconfig "social_pay/config"
	mylibs "social_pay/libs"
	mymodel "social_pay/model"
	mytypedef "social_pay/typedef"
	"strconv"

	"github.com/time2k/letsgo-ng"
)

//PRPC 结构体
type PRPC struct {
	Commp  letsgo.CommonParams
	Logger *log.Logger
}

//CreateOrder 支付下单接口
func (prpc *PRPC) CreateOrder(orderParams mytypedef.OrderCreateParams, res *mytypedef.OrderResponse) error {
	prpc.Logger.Println("call PRPC.CreateOrder")
	//panic recover
	defer letsgo.PanicFunc()

	//非空参数校验
	if orderParams.BusinessId == 0 {
		return fmt.Errorf("%d %s", myconfig.StatusError, "business_id参数不能为空")
	}
	if orderParams.ProductId == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "product_id参数不能为空")
	}
	if orderParams.ProductName == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "product_name参数不能为空")
	}
	if orderParams.OutTradeNo == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "out_trade_no参数不能为空")
	}
	if orderParams.TotalFee == 0 {
		return fmt.Errorf("%d %s", myconfig.StatusError, "total_fee参数不能为空")
	}
	if orderParams.NotifyUrl == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "notify_url参数不能为空")
	}
	if orderParams.ClientIp == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "client_ip参数不能为空")
	}
	if orderParams.SignType == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "sign_type参数不能为空")
	}
	if orderParams.Sign == "" {
		return fmt.Errorf("%d %s", myconfig.StatusError, "sign参数不能为空")
	}
	//校验支付方式
	business_id := orderParams.BusinessId
	ret := mymodel.GetBusinessInfoList(letsgo.CommonParams{})
	if ret.Status != myconfig.StatusOk {
		return fmt.Errorf("%d %s", myconfig.StatusError, "业务配置不存在")
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
		return fmt.Errorf("%d %s", myconfig.StatusError, "未查询到对应business_id")
	}

	//公众号&小程序支付必需open_id
	if business_info.BusinessPayChannel == "WXPAY_PUB" {
		if orderParams.OpenId == "" {
			return fmt.Errorf("%d %s", myconfig.StatusError, "open_id参数不能为空")
		}
	}

	//校验业务方订单号是否重复下单
	checkOutNo := mymodel.CheckOutTradeNo(prpc.Commp, orderParams.OutTradeNo)
	if checkOutNo.Status == myconfig.StatusOk && checkOutNo.Body == true {
		return fmt.Errorf("%d %s", myconfig.StatusError, "out_trade_no订单号重复下单")
	}
	//校验业务线签名是否正确
	checkSign := mymodel.CheckOrderSign(prpc.Commp, &orderParams)
	if checkSign.Status == myconfig.StatusOk && checkSign.Body == false {
		return fmt.Errorf("%d %s", myconfig.StatusError, "sign加密签名错误")
	}
	//生成支付订单号 存储订单信息
	insertRes := mymodel.InsertPayOrder(prpc.Commp, &orderParams)
	if insertRes.Status != myconfig.StatusOk {
		return fmt.Errorf("%d %s", myconfig.StatusError, insertRes.Msg)
	}
	orderId := insertRes.Body.(int64)

	switch business_info.BusinessPayChannel {
	case "ALIPAY_H5":
		resData := mymodel.AliH5PayOrder(prpc.Commp, orderId, orderParams.BusinessId, orderParams.ReturnUrl)
		*res = resData.Body.(mytypedef.OrderResponse)
	case "WXPAY_H5":
		resData := mymodel.WxH5PayOrder(prpc.Commp, orderId, orderParams.BusinessId)
		*res = resData.Body.(mytypedef.OrderResponse)
	case "WXPAY_PUB":
		resData := mymodel.WxPublicPayOrder(prpc.Commp, orderId, orderParams.BusinessId, orderParams.OpenId)
		*res = resData.Body.(mytypedef.OrderResponse)
	case "ALPAY_APP":
		resData := mymodel.AliAppPayOrder(prpc.Commp, orderId, orderParams.BusinessId)
		*res = resData.Body.(mytypedef.OrderResponse)
	case "WXPAY_APP":
		resData := mymodel.WxAppPayOrder(prpc.Commp, orderId, orderParams.BusinessId)
		*res = resData.Body.(mytypedef.OrderResponse)
	case "APPLEPAY":
		resData := mymodel.IapPayOrder(prpc.Commp, orderId)
		*res = resData.Body.(mytypedef.OrderResponse)
	}

	return nil
}

//IapPayVerify 苹果IAP支付结果验证
func (prpc *PRPC) IapPayVerify(verifyParam mytypedef.IAPVerifyParam, res *mytypedef.OrderResponse) error {
	prpc.Logger.Println("call PRPC.IapPayVerify")
	//panic recover
	defer letsgo.PanicFunc()

	receiptData := verifyParam.ReceiptData                                     //验证数据
	useSandbox, err := strconv.ParseBool(strconv.Itoa(verifyParam.UseSandbox)) //是否沙盒环境
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
	orderNo := verifyParam.OrderNo
	if orderNo == "" {
		return fmt.Errorf("%d %s", myconfig.StatusParamsNoValid, "订单号不能为空")
	}
	payOrderInfo := mymodel.PayOrderInfoByNo(prpc.Commp, orderNo)
	if payOrderInfo.Status == myconfig.StatusOk {
		log.Println("CALLBACK IAP", payOrderInfo)
		payOrderData := payOrderInfo.Body.(*mytypedef.PayOrder)
		//todo 此处暂验证苹果充值productId = 下单时的productId
		if receipt.InApp[0].ProductId == payOrderData.ProductId {
			log.Println("CALLBACK ProductId", receipt.InApp[0])
			updateOrder := mymodel.PayOrderInfoUpdate(prpc.Commp, orderNo, 1, 1, useSandbox)
			if updateOrder.Status == myconfig.StatusOk {
				//todo 异步通知第三方支付成功结果
				log.Println("CALLBACK CallBackBusinessResult", orderNo)
				go mymodel.CallBackBusinessResult(prpc.Commp, orderNo)
				//todo 异步通知第三方支付成功结果
				res.OrderNo = payOrderData.OrderNo
				res.ReturnType = "IAP_PAY"
				res.ReturnData = receipt
			} else {
				log.Println("IAP NOTIFY ORDER UPDATE FAILED", orderNo)
			}
		}
	}

	return nil
}
