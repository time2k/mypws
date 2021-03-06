package mylibs

type WeChatUnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	TradeType  string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	PrepayId   string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty"`
	CodeUrl    string `xml:"code_url,omitempty" json:"code_url,omitempty"`
	MwebUrl    string `xml:"mweb_url,omitempty" json:"mweb_url,omitempty"`
	Timestamp  string `xml:"timestamp,omitempty" json:"timestamp,omitempty"`
}

type WeChatQueryOrderResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState         string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        int    `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponFee0         int    `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         int    `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc     string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
}

type WeChatCloseOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}

type WeChatReverseResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Recall     string `xml:"recall,omitempty" json:"recall,omitempty"`
}

type WeChatRefundResponse struct {
	ReturnCode          string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg           string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode          string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid               string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId               string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty" json:"sign,omitempty"`
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	RefundFee           int    `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	TotalFee            int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee             int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashRefundFee       int    `xml:"cash_refund_fee,omitempty" json:"cash_refund_fee,omitempty"`
	CouponType0         string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1         string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponRefundFee     int    `xml:"coupon_refund_fee,omitempty" json:"coupon_refund_fee,omitempty"`
	CouponRefundFee0    int    `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1    int    `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CouponRefundCount   int    `xml:"coupon_refund_count,omitempty" json:"coupon_refund_count,omitempty"`
	CouponRefundId0     string `xml:"coupon_refund_id_0,omitempty" json:"coupon_refund_id_0,omitempty"`
	CouponRefundId1     string `xml:"coupon_refund_id_1,omitempty" json:"coupon_refund_id_1,omitempty"`
}

type WeChatQueryRefundResponse struct {
	ReturnCode           string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg            string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode           string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode              string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes           string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid                string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId                string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr             string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                 string `xml:"sign,omitempty" json:"sign,omitempty"`
	TotalRefundCount     int    `xml:"total_refund_count,omitempty" json:"total_refund_count,omitempty"`
	TransactionId        string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo           string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee             int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee   int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType              string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee              int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	RefundCount          int    `xml:"refund_count,omitempty" json:"refund_count,omitempty"`
	OutRefundNo0         string `xml:"out_refund_no_0,omitempty" json:"out_refund_no_0,omitempty"`
	OutRefundNo1         string `xml:"out_refund_no_1,omitempty" json:"out_refund_no_1,omitempty"`
	RefundId0            string `xml:"refund_id_0,omitempty" json:"refund_id_0,omitempty"`
	RefundId1            string `xml:"refund_id_1,omitempty" json:"refund_id_1,omitempty"`
	RefundChannel0       string `xml:"refund_channel_0,omitempty" json:"refund_channel_0,omitempty"`
	RefundChannel1       string `xml:"refund_channel_1,omitempty" json:"refund_channel_1,omitempty"`
	RefundFee            int    `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	RefundFee0           int    `xml:"refund_fee_0,omitempty" json:"refund_fee_0,omitempty"`
	RefundFee1           int    `xml:"refund_fee_1,omitempty" json:"refund_fee_1,omitempty"`
	SettlementRefundFee0 int    `xml:"settlement_refund_fee_0,omitempty" json:"settlement_refund_fee_0,omitempty"`
	SettlementRefundFee1 int    `xml:"settlement_refund_fee_1,omitempty" json:"settlement_refund_fee_1,omitempty"`
	CouponType00         string `xml:"coupon_type_0_0,omitempty" json:"coupon_type_0_0,omitempty"`
	CouponType01         string `xml:"coupon_type_0_1,omitempty" json:"coupon_type_0_1,omitempty"`
	CouponType10         string `xml:"coupon_type_1_0,omitempty" json:"coupon_type_1_0,omitempty"`
	CouponType11         string `xml:"coupon_type_1_1,omitempty" json:"coupon_type_1_1,omitempty"`
	CouponRefundFee0     int    `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1     int    `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CouponRefundCount0   int    `xml:"coupon_refund_count_0,omitempty" json:"coupon_refund_count_0,omitempty"`
	CouponRefundCount1   int    `xml:"coupon_refund_count_1,omitempty" json:"coupon_refund_count_1,omitempty"`
	CouponRefundId00     string `xml:"coupon_refund_id_0_0,omitempty" json:"coupon_refund_id_0_0,omitempty"`
	CouponRefundId01     string `xml:"coupon_refund_id_0_1,omitempty" json:"coupon_refund_id_0_1,omitempty"`
	CouponRefundId10     string `xml:"coupon_refund_id_1_0,omitempty" json:"coupon_refund_id_1_0,omitempty"`
	CouponRefundId11     string `xml:"coupon_refund_id_1_1,omitempty" json:"coupon_refund_id_1_1,omitempty"`
	CouponRefundFee00    int    `xml:"coupon_refund_fee_0_0,omitempty" json:"coupon_refund_fee_0_0,omitempty"`
	CouponRefundFee01    int    `xml:"coupon_refund_fee_0_1,omitempty" json:"coupon_refund_fee_0_1,omitempty"`
	CouponRefundFee10    int    `xml:"coupon_refund_fee_1_0,omitempty" json:"coupon_refund_fee_1_0,omitempty"`
	CouponRefundFee11    int    `xml:"coupon_refund_fee_1_1,omitempty" json:"coupon_refund_fee_1_1,omitempty"`
	RefundStatus0        string `xml:"refund_status_0,omitempty" json:"refund_status_0,omitempty"`
	RefundStatus1        string `xml:"refund_status_1,omitempty" json:"refund_status_1,omitempty"`
	RefundAccount0       string `xml:"refund_account_0,omitempty" json:"refund_account_0,omitempty"`
	RefundAccount1       string `xml:"refund_account_1,omitempty" json:"refund_account_1,omitempty"`
	RefundRecvAccout0    string `xml:"refund_recv_accout_0,omitempty" json:"refund_recv_accout_0,omitempty"`
	RefundRecvAccout1    string `xml:"refund_recv_accout_1,omitempty" json:"refund_recv_accout_1,omitempty"`
	RefundSuccessTime0   string `xml:"refund_success_time_0,omitempty" json:"refund_success_time_0,omitempty"`
	RefundSuccessTime1   string `xml:"refund_success_time_1,omitempty" json:"refund_success_time_1,omitempty"`
}

type WeChatMicropayResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	PromotionDetail    string `xml:"promotion_detail,omitempty" json:"promotion_detail,omitempty"`
}

type WeChatTransfersResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	MchAppid       string `xml:"mch_appid,omitempty" json:"mch_appid,omitempty"`
	Mchid          string `xml:"mchid,omitempty" json:"mchid,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty" json:"payment_no,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty" json:"payment_time,omitempty"`
}

type getSignKeyResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SandboxSignkey string `xml:"sandbox_signkey,omitempty" json:"sandbox_signkey,omitempty"`
}

type WeChatNotifyRequest struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	SignType           string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        int    `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponFee0         int    `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         int    `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
}

type WeChatRefundNotifyRequest struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ReqInfo    string `xml:"req_info,omitempty" json:"req_info,omitempty"`
}

type RefundNotify struct {
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	TotalFee            int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	RefundFee           int    `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	RefundStatus        string `xml:"refund_status,omitempty" json:"refund_status,omitempty"`
	SuccessTime         string `xml:"success_time,omitempty" json:"success_time,omitempty"`
	RefundRecvAccout    string `xml:"refund_recv_accout,omitempty" json:"refund_recv_accout,omitempty"`
	RefundAccount       string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`
	RefundRequestSource string `xml:"refund_request_source,omitempty" json:"refund_request_source,omitempty"`
}

type Code2SessionRsp struct {
	SessionKey string `json:"session_key,omitempty"` // ????????????
	ExpiresIn  int    `json:"expires_in,omitempty"`  // SessionKey?????????????????????
	Openid     string `json:"openid,omitempty"`      // ??????????????????
	Unionid    string `json:"unionid,omitempty"`     // ???????????????????????????????????????
	Errcode    int    `json:"errcode,omitempty"`     // ?????????
	Errmsg     string `json:"errmsg,omitempty"`      // ????????????
}

type PaidUnionId struct {
	Unionid string `json:"unionid,omitempty"` // ???????????????????????????????????????
	Errcode int    `json:"errcode,omitempty"` // ?????????
	Errmsg  string `json:"errmsg,omitempty"`  // ????????????
}

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // ??????????????????
	ExpiresIn   int    `json:"expires_in,omitempty"`   // SessionKey?????????????????????
	Errcode     int    `json:"errcode,omitempty"`      // ?????????
	Errmsg      string `json:"errmsg,omitempty"`       // ????????????
}

type WeChatUserInfo struct {
	Subscribe      int    `json:"subscribe,omitempty"`       // ?????????????????????????????????????????????0???????????????????????????????????????????????????????????????????????????
	Openid         string `json:"openid,omitempty"`          // ??????????????????
	Nickname       string `json:"nickname,omitempty"`        // ???????????????
	Sex            int    `json:"sex,omitempty"`             // ????????????????????????1?????????????????????2?????????????????????0????????????
	Language       string `json:"language,omitempty"`        // ?????????????????????????????????zh_CN
	City           string `json:"city,omitempty"`            // ??????????????????
	Province       string `json:"province,omitempty"`        // ??????????????????
	Country        string `json:"country,omitempty"`         // ??????????????????
	Headimgurl     string `json:"headimgurl,omitempty"`      // ??????????????????????????????????????????????????????????????????0???46???64???96???132???????????????0??????640*640?????????????????????????????????????????????????????????????????????????????????????????????URL????????????
	SubscribeTime  int    `json:"subscribe_time,omitempty"`  // ??????????????????????????????????????????????????????????????????????????????????????????
	Unionid        string `json:"unionid,omitempty"`         // ??????????????????????????????????????????????????????????????????????????????????????????
	Remark         string `json:"remark,omitempty"`          // ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
	Groupid        int    `json:"groupid,omitempty"`         // ?????????????????????ID????????????????????????????????????
	TagidList      []int  `json:"tagid_list,omitempty"`      // ????????????????????????ID??????
	SubscribeScene string `json:"subscribe_scene,omitempty"` // ????????????????????????????????????ADD_SCENE_SEARCH ??????????????????ADD_SCENE_ACCOUNT_MIGRATION ??????????????????ADD_SCENE_PROFILE_CARD ???????????????ADD_SCENE_QR_CODE ??????????????????ADD_SCENEPROFILE LINK ???????????????????????????ADD_SCENE_PROFILE_ITEM ???????????????????????????ADD_SCENE_PAID ??????????????????ADD_SCENE_OTHERS ??????
	QrScene        int    `json:"qr_scene,omitempty"`        // ?????????????????????????????????????????????
	QrSceneStr     string `json:"qr_scene_str,omitempty"`    // ???????????????????????????????????????????????????
	Errcode        int    `json:"errcode,omitempty"`         // ?????????
	Errmsg         string `json:"errmsg,omitempty"`          // ????????????
}

// ???????????????????????? ????????????????????????
type WeChatUserPhone struct {
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	PurePhoneNumber string         `json:"purePhoneNumber,omitempty"`
	CountryCode     string         `json:"countryCode,omitempty"`
	Watermark       *watermarkInfo `json:"watermark,omitempty"`
}

// ???????????????????????? ?????????????????????
type WeChatAppletUserInfo struct {
	OpenId    string         `json:"openId,omitempty"`
	NickName  string         `json:"nickName,omitempty"`
	Gender    int            `json:"gender,omitempty"`
	City      string         `json:"city,omitempty"`
	Province  string         `json:"province,omitempty"`
	Country   string         `json:"country,omitempty"`
	AvatarUrl string         `json:"avatarUrl,omitempty"`
	UnionId   string         `json:"unionId,omitempty"`
	Watermark *watermarkInfo `json:"watermark,omitempty"`
}

type watermarkInfo struct {
	Appid     string `json:"appid,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

// ???????????????openid ??????
type OpenIdByAuthCodeRsp struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	Openid     string `xml:"openid,omitempty" json:"openid,omitempty"` // ??????????????????
}

// App??????????????????????????????code??????access_token
type AppWeChatLoginAccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	Openid       string `json:"openid,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Unionid      string `json:"unionid,omitempty"`
	Errcode      int    `json:"errcode,omitempty"` // ?????????
	Errmsg       string `json:"errmsg,omitempty"`  // ????????????
}

// ??????App?????????????????????????????????????????? access_token
type RefreshAppWeChatLoginAccessTokenRsp struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	Openid       string `json:"openid,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Errcode      int    `json:"errcode,omitempty"` // ?????????
	Errmsg       string `json:"errmsg,omitempty"`  // ????????????
}
