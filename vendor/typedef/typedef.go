package mytypedef

//Empty 结构体
type Empty struct {
}

type OrderCreateParams struct {
	BusinessId  int     `json:"business_id" form:"business_id" validate:"required"`
	BusinessUid string  `json:"business_uid" form:"business_uid" validate:""`
	ProductId   string  `json:"product_id" form:"product_id" validate:"required"`
	ProductName string  `json:"product_name" form:"product_name" validate:"required"`
	ProductDesc string  `json:"product_desc" form:"product_desc" validate:""`
	OutTradeNo  string  `json:"out_trade_no" form:"out_trade_no" validate:"required"`
	TotalFee    float64 `json:"total_fee" form:"total_fee" validate:"required"`
	NotifyUrl   string  `json:"notify_url" form:"notify_url" validate:"required,url"`
	ClientIp    string  `json:"client_ip" form:"client_ip" validate:"required"`
	OpenId      string  `json:"open_id" form:"open_id" validate:""`
	ReturnUrl   string  `json:"return_url" form:"return_url" validate:"url"`
	SignType    string  `json:"sign_type" form:"sign_type" validate:"required"`
	Sign        string  `json:"sign" form:"sign" validate:"required"`
}

type PayOrder struct {
	Id           int64   `json:"id"`
	OrderNo      string  `json:"order_no"`
	BusinessId   int64   `json:"business_id"`
	BusinessName string  `json:"business_name"`
	BusinessUid  string  `json:"business_uid"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductDesc  string  `json:"product_desc"`
	OutTradeNo   string  `json:"out_trade_no"`
	TotalFee     float64 `json:"total_fee"`
	OrderStatus  int     `json:"order_status"`
	NotifyStatus int     `json:"notify_status"`
	NotifyUrl    string  `json:"notify_url"`
	ClientIp     string  `json:"client_ip"`
	ServerIp     string  `json:"server_ip"`
	Sandbox      int     `json:"sandbox"`
	CreateTime   int64   `json:"create_time"`
	UpdateTime   int64   `json:"update_time"`
}

type BusinessPayInfo struct {
	Id                 int64  `json:"id"`
	BusinessId         int64  `json:"business_id"`
	BusinessPayChannel string `json:"business_paychannel"`
	BusinessName       string `json:"business_name"`
	BusinessKey        string `json:"business_key"`
	AlipayAppId        string `json:"alipay_app_id"`
	AlipayProductCode  string `json:"alipay_product_code"`
	AlipayNotifyUrl    string `json:"alipay_notify_url"`
	AlipayPrivateKey   string `json:"alipay_private_key"`
	AlipayPublicKey    string `json:"alipay_public_key"`
	WxAppId            string `json:"wx_app_id"`
	WxMchId            string `json:"wx_mch_id"`
	WxTradeType        string `json:"wx_trade_type"`
	WxSignKey          string `json:"wx_sign_key"`
	WxNotifyUrl        string `json:"wx_notify_url"`
	WxBody             string `json:"wx_body"`
	Status             int    `json:"status"`
	CreateTime         int64  `json:"create_time"`
	UpdateTime         int64  `json:"update_time"`
}

type BusinessPayInfoList struct {
	List []BusinessPayInfo `json:"list"`
}

type FailedNotifyOrder struct {
	Id          int64  `json:"id"`
	OrderNo     string `json:"order_no"`
	Type        int    `json:"type"`
	Status      int    `json:"status"`
	FailedCount int    `json:"failed_count"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

type FailedNotifyOrderList struct {
	List []FailedNotifyOrder `json:"list"`
}

type OrderResponse struct {
	ReturnType string      `json:"return_type"`
	ReturnData interface{} `json:"return_data"`
	OrderNo    string      `json:"order_no"`
}

type IAPVerifyParam struct {
	ReceiptData string `json:"receipt_data"`
	UseSandbox  int    `json:"use_sandbox"`
	OrderNo     string `json:"order_no"`
}
