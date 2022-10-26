package unionpay

// 常量参数
const (
	Format      = "JSON"
	Charset     = "UTF-8"
	ContentType = "application/json"
)

// CommonResponseParams 通用应答参数，用于判断网关状态码
type CommonResponseParams struct {
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	ErrInfo string `json:"errInfo"`
}

// RefundRequest 交易退款请求
type RefundRequest struct {
	MsgId            string `json:"msgId"`            //消息ID
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	MerOrderId       string `json:"merOrderId"`       // 必填 商户单号
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	Mid              string `json:"mid"`              //商户号
	Tid              string `json:"tid"`              // 终端号
	SubAppId         string `json:"subAppId"`         //条件必填 微信子商 户appId
	InstMid          string `json:"instMid"`          // 业务类型 APPDEFAULT
	RefundAmount     int64  `json:"refundAmount"`     // 要退货的金额 若下单接口中上送了分账标记字段divisionFlag，则该字段 refundAmount=subOrders 中totalAmount之和 +platformAmount
	PlatformAmount   int64  `json:"platformAmount"`   //平台商户退款分账金额, 若原交易是分账交易，则分账金额必传，且退款接口 platformAmount小于下单接口中上送的platformAmount
	RefundDesc       string `json:"refundDesc"`       //退货说明
}

// RefundResponse 交易退款应答
type RefundResponse struct {
	MsgId               string `json:"msgId"` //非必填 消息ID， 原样返回
	ErrCode             string `json:"errCode"`
	ErrMsg              string `json:"errMsg"`
	ResponseTimeStamp   string `json:"responseTimeStamp"`   // 必填 报文应答时间 yyyy-MM-dd HH:mm:ss
	SrcReserve          string `json:"srcReserve"`          //非必填 请求系统 预留字段
	Mid                 string `json:"mid"`                 //商户号
	Tid                 string `json:"tid"`                 // 终端号
	MerOrderId          string `json:"merOrderId"`          // 必填 商户单号
	MerName             string `json:"merName"`             //商户名称
	SeqId               string `json:"seqId"`               //平台流水号
	Status              string `json:"status"`              //交易状态
	TargetMid           string `json:"targetMid"`           //支付渠道商户 号
	TargetOrderId       string `json:"targetOrderId"`       //目标平台单号
	TargetStatus        string `json:"targetStatus"`        //目标平台状态
	TargetSys           string `json:"targetSys"`           //目标平台代码
	TotalAmount         string `json:"totalAmount"`         //支付总金额
	RefundAmount        string `json:"refundAmount"`        //总退款金额
	RefundFunds         string `json:"refundFunds"`         //退款渠道列表
	RefundFundsDesc     string `json:"refundFundsDesc"`     //退款渠道描述
	RefundOrderId       string `json:"refundOrderId"`       //退货订单号
	RefundTargetOrderId string `json:"refundTargetOrderId"` //目标系统退货订单号
	RefundInvoiceAmount string `json:"refundInvoiceAmount"` //实付部分退款金额
	YxlmAmount          string `json:"yxlmAmount"`          //营销联盟优惠金额
	RefundStatus        string `json:"refundStatus"`        //退款状态
}

// QueryRefundRequest 交易退款查询请求
type QueryRefundRequest struct {
	MsgId            string `json:"msgId"`            //非必填 消息ID， 原样返回
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	MerOrderId       string `json:"merOrderId"`       // 必填 商户单号
	InstMid          string `json:"instMid"`          // 业务类型 APPDEFAULT
	Mid              string `json:"mid"`              // 必填 商户号
	Tid              string `json:"tid"`              // 必填 终端号
}

// QueryRefundResponse 交易退款查询应答
type QueryRefundResponse struct {
	ErrCode             string `json:"errCode"`
	ErrMsg              string `json:"errMsg"`
	MsgId               string `json:"msgId"`               //非必填 消息ID， 原样返回
	ResponseTimeStamp   string `json:"responseTimeStamp"`   // 必填 报文应答时间 yyyy-MM-dd HH:mm:ss
	SrcReserve          string `json:"srcReserve"`          //非必填 请求系统 预留字段
	RefundStatus        string `json:"refundStatus"`        // 退款状态
	RefundOrderId       string `json:"refundOrderId"`       //退货订单号
	RefundTargetOrderId string `json:"refundTargetOrderId"` //目标系统退货订单号
	Mid                 string `json:"mid"`                 //商户号
	Tid                 string `json:"tid"`                 // 终端号
	SeqId               string `json:"seqId"`               //平台流水号
	SettleRefId         string `json:"settleRefId"`         //清分ID 如果来源方传了 bankRefId就等 于bankRefId， 否则等于seqId
	Status              string `json:"status"`              //交易状态
	TotalAmount         string `json:"totalAmount"`         //支付总金额
	MerName             string `json:"merName"`             //商户名称
	MerOrderId          string `json:"merOrderId"`          //商户订单号
	TargetOrderId       string `json:"targetOrderId"`       //目标平台单号
	TargetSys           string `json:"targetSys"`           //目标平台代码
	TargetStatus        string `json:"targetStatus"`        //目标平台状态
	TargetMid           string `json:"targetMid"`           //支付渠道商户 号
	BankCardNo          string `json:"bankCardNo"`
	BankInfo            string `json:"bankInfo"`
	RefundAmount        string `json:"refundAmount"`        //总退款金额
	RefundFunds         string `json:"refundFunds"`         //退款渠道列表
	RefundFundsDesc     string `json:"refundFundsDesc"`     //退款渠道描述
	PayTime             string `json:"payTime"`             //支付时间
	SettleDate          string `json:"settleDate"`          //结算日期
	RefundInvoiceAmount string `json:"refundInvoiceAmount"` //实付部分退款金额
	YxlmAmount          string `json:"yxlmAmount"`          //营销联盟优惠金额
	SendBackAmount      string `json:"sendBackAmount"`      //商户实退金额
}

// URL 接口地址
const (
	AliPreOrderApiBeta             = "https://test-api-open.chinaums.com/v1/netpay/trade/precreate"     // 下单接口-支付宝
	WxPreOrderApiBeta              = "https://test-api-open.chinaums.com/v1/netpay/wx/app-pre-order"    // 下单接口-微信
	QmfPreOrderApiBeta             = "https://test-api-open.chinaums.com/v1/netpay/qmf/order"           // 下单接口-全民支付
	UacPreOrderApiBeta             = "https://test-api-open.chinaums.com/v1/netpay/uac/app-order"       // 下单接口-银联云闪付
	ApplePreOrderApiBeta           = "https://test-api-open.chinaums.com/v1/netpay/applepay/order"      // 下单接口-Apple Pay
	AliMiniAppPreOrderApiBeta      = "https://test-api-open.chinaums.com/v1/netpay/trade/app-pre-order" // 下单接口-支付宝(支付宝跳转小程序支付)
	OrderQueryApiBeta              = "https://test-api-open.chinaums.com/v1/netpay/query"               //订单交易查询
	WxMiniAppOrderQueryApiBeta     = "https://test-api-open.chinaums.com/v1/netpay/wx/app-pre-query"    //订单交易查询-微信(微信跳转小程序专用)
	AliMiniAppOrderQueryApiBeta    = "https://test-api-open.chinaums.com/v1/netpay/trade/app-pre-query" //订单交易查询-支付宝(支付宝跳转小程序专用)
	PreOrderCloseApiBeta           = "https://test-api-open.chinaums.com/v1/netpay/close"               // 订单关闭-用户创建订单之后，对未支付的订单进行关闭操作
	WxMiniAppPreOrderCloseApiBeta  = "https://test-api-open.chinaums.com/v1/netpay/wx/app-pre-close"    // 订单关闭-微信(微信跳转小程序专用)
	AliMiniAppPreOrderCloseApiBeta = "https://test-api-open.chinaums.com/v1/netpay/trade/app-pre-close" // 订单关闭-支付宝(支付宝跳转小程序支付)
	RefundApiBeta                  = "https://test-api-open.chinaums.com/v1/netpay/refund"              //退款
	RefundQueryApiBeta             = "https://test-api-open.chinaums.com/v1/netpay/refund-query"        //退款查询

	AliPreOrderApi             = "https://api-mop.chinaums.com/v1/netpay/trade/precreate"     // 下单接口-支付宝
	WxPreOrderApi              = "https://api-mop.chinaums.com/v1/netpay/wx/app-pre-order"    // 下单接口-微信
	QmfPreOrderApi             = "https://api-mop.chinaums.com/v1/netpay/qmf/order"           // 下单接口-全民支付
	UacPreOrderApi             = "https://api-mop.chinaums.com/v1/netpay/uac/app-order"       // 下单接口-银联云闪付
	ApplePreOrderApi           = "https://api-mop.chinaums.com/v1/netpay/applepay/order"      // 下单接口-Apple Pay
	AliMiniAppPreOrderApi      = "https://api-mop.chinaums.com/v1/netpay/trade/app-pre-order" // 下单接口-支付宝(支付宝跳转小程序支付)
	OrderQueryApi              = "https://api-mop.chinaums.com/v1/netpay/query"               //订单交易查询
	WxMiniAppOrderQueryApi     = "https://api-mop.chinaums.com/v1/netpay/wx/app-pre-query"    //订单交易查询-微信(微信跳转小程序专用)
	AliMiniAppOrderQueryApi    = "https://api-mop.chinaums.com/v1/netpay/trade/app-pre-query" //订单交易查询-支付宝(支付宝跳转小程序专用)
	PreOrderCloseApi           = "https://api-mop.chinaums.com/v1/netpay/close"               // 订单关闭-用户创建订单之后，对未支付的订单进行关闭操作
	WxMiniAppPreOrderCloseApi  = "https://api-mop.chinaums.com/v1/netpay/wx/app-pre-close"    // 订单关闭-微信(微信跳转小程序专用)
	AliMiniAppPreOrderCloseApi = "https://api-mop.chinaums.com/v1/netpay/trade/app-pre-close" // 订单关闭-支付宝(支付宝跳转小程序支付)
	RefundApi                  = "https://api-mop.chinaums.com/v1/netpay/refund"              //退款
	RefundQueryApi             = "https://api-mop.chinaums.com/v1/netpay/refund-query"        //退款查询
)

const (
	GateWaySuccess = "SUCCESS" // 网关成功状态码
)

//PreOrderRequestCommon 公共支付请求参数
type PreOrderRequestCommon struct {
	MsgId            string `json:"msgId"`            //非必填 消息ID， 原样返回
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	MerOrderId       string `json:"merOrderId"`       // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	Mid              string `json:"mid"`              //商户号
	Tid              string `json:"tid"`              // 终端号
	SubAppId         string `json:"subAppId"`         //条件必填 微信子商 户appId
	InstMid          string `json:"instMid"`          // 业务类型
	TotalAmount      int64  `json:"totalAmount"`      // 支付总金额
	TradeType        string `json:"tradeType"`        // APP 交易类型，微信必传
	NotifyUrl        string `json:"notifyUrl"`        //非必填 支付结果 通知地址
	ShowUrl          string `json:"showUrl"`          //非必填 订单展示 页面
}

//AliPreOrderRequest 支付宝支付请求参数
type AliPreOrderRequest struct {
	MsgId            string `json:"msgId"`            //非必填 消息ID， 原样返回
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	MerOrderId       string `json:"merOrderId"`       // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	Mid              string `json:"mid"`              //商户号
	Tid              string `json:"tid"`              // 终端号
	SubAppId         string `json:"subAppId"`         //条件必填 微信子商 户appId
	InstMid          string `json:"instMid"`          // 业务类型
	TotalAmount      int64  `json:"totalAmount"`      // 支付总金额
	TradeType        string `json:"tradeType"`        // APP 交易类型，微信必传
	NotifyUrl        string `json:"notifyUrl"`        //非必填 支付结果 通知地址
	ShowUrl          string `json:"showUrl"`          //非必填 订单展示 页面
}

//WxPreOrderRequest 微信支付请求参数
type WxPreOrderRequest struct {
	MsgId            string `json:"msgId"`            //非必填 消息ID， 原样返回
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	MerOrderId       string `json:"merOrderId"`       // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	Mid              string `json:"mid"`              //商户号
	Tid              string `json:"tid"`              // 终端号
	SubAppId         string `json:"subAppId"`         //条件必填 微信子商 户appId
	InstMid          string `json:"instMid"`          // 业务类型
	TotalAmount      int64  `json:"totalAmount"`      // 支付总金额
	TradeType        string `json:"tradeType"`        // APP 交易类型，微信必传
	NotifyUrl        string `json:"notifyUrl"`        //非必填 支付结果 通知地址
	ShowUrl          string `json:"showUrl"`          //非必填 订单展示 页面
}

//PreOrderResponse app支付响应参数
type PreOrderResponse struct {
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

//ClosePreOrderRequest 关闭订单请求参数
type ClosePreOrderRequest struct {
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	MerOrderId       string `json:"merOrderId"`       // 条件必填 商户单号，全局唯一，不可重复，长度不超过50位
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	Mid              string `json:"mid"`              //商户号
	Tid              string `json:"tid"`              // 终端号
	InstMid          string `json:"instMid"`          // 业务类型
}

//ClosePreOrderResponse 关闭订单响应参数
type ClosePreOrderResponse struct {
	ResponseTimestamp string `json:"responseTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	Mid               string `json:"mid"`               //商户号
	Tid               string `json:"tid"`               // 终端号
	InstMid           string `json:"instMid"`           // 业务类型
	ErrCode           string `json:"errCode"`
	ErrMsg            string `json:"errMsg"`
}

//QueryOrderRequest 查询交易订单请求参数
type QueryOrderRequest struct {
	RequestTimestamp string `json:"requestTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	SrcReserve       string `json:"srcReserve"`       //非必填 请求系统 预留字段
	Mid              string `json:"mid"`              //商户号
	Tid              string `json:"tid"`              // 终端号
	InstMid          string `json:"instMid"`          // 业务类型 APPDEFAULT
	MerOrderId       string `json:"merOrderId"`       // 必填 商户单号，全局唯一，不可重复，长度不超过50位
	TargetOrderId    string `json:"targetOrderId"`    //必填 支付订单号，全局唯一，不可重复，长度不超过50位
}

//QueryOrderResponse 查询交易订单响应参数
type QueryOrderResponse struct {
	ResponseTimestamp string `json:"responseTimestamp"` // 必填 报文请求时间 yyyy-MM-dd HH:mm:ss
	Mid               string `json:"mid"`               //商户号
	Tid               string `json:"tid"`               // 终端号
	InstMid           string `json:"instMid"`           // 业务类型 APPDEFAULT
	ErrCode           string `json:"errCode"`
	ErrMsg            string `json:"errMsg"`
}
