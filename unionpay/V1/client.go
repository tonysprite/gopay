package unionpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay/pkg/xlog"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type BaseClient struct {
	RequestId  string
	HttpClient *http.Client

	RequestBody  []byte
	ResponseBody []byte
}

func NewBaseClient(requestId string) *BaseClient {
	c := new(BaseClient)
	c.HttpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 90 * time.Second,
				// DualStack: true,
			}).DialContext,
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
			// ExpectContinueTimeout: 1 * time.Second,
			// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 40 * time.Second, // 整个http请求发起到等待应答的超时时间
	}
	c.RequestId = requestId
	return c
}

func (r *BaseClient) SetRequestBody(body []byte) {
	body = r.TrimSpace(body)
	r.RequestBody = body
}

func (r *BaseClient) SetResponseBody(body []byte) {
	body = r.TrimSpace(body)
	r.ResponseBody = body
}

func (r *BaseClient) GetRequestBody() string {
	return string(r.RequestBody)
}

func (r *BaseClient) GetResponseBody() string {
	return string(r.ResponseBody)
}

func (r *BaseClient) TrimSpace(body []byte) []byte {
	body = bytes.TrimSpace(body)
	body = bytes.Replace(body, []byte("\r"), []byte(""), -1)
	body = bytes.Replace(body, []byte("\n"), []byte(""), -1)
	return body
}

func (r *BaseClient) GetGateWayError(gatewayCode, gatewayMsg string) error {
	return fmt.Errorf(r.GetGateWayStatus(gatewayCode, gatewayMsg))
}

func (r *BaseClient) GetGateWayStatus(gatewayCode, gatewayMsg string) string {
	return fmt.Sprintf("gatewayCode:%s gatewayMsg:%s", gatewayCode, gatewayMsg)
}

func (r *BaseClient) GetBizError(bizCode, bizMsg string) error {
	return fmt.Errorf(r.GetBizTrxMsg(bizCode, bizMsg))
}

func (r *BaseClient) GetBizTrxMsg(bizCode, bizMsg string) string {
	return fmt.Sprintf("bizCode:%s bizMsg:%s", bizCode, bizMsg)
}

type Client struct {
	*BaseClient

	// 服务商配置信息 商户级别的
	AppId        string // 应用ID 实际交易的商户号
	AppKey       string // 签名秘钥
	MerchantCode string // 服务商商户号
	IsProEnv     bool   // 是否生产环境
}

func NewClient(requestId, appId, appKey string, isProEnv bool) *Client {
	c := &Client{
		BaseClient: NewBaseClient(requestId),
		AppId:      appId,
		AppKey:     appKey,
		IsProEnv:   isProEnv,
	}
	return c
}

// 请求API地址
func (c *Client) requestApi(reqObj interface{}, addr string) ([]byte, error) {
	body, err := json.Marshal(reqObj)
	if err != nil {
		return nil, err
	}

	c.SetRequestBody(body)

	// 发起请求
	req, err := http.NewRequest("POST", addr, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	authorization, err := Sign(body, c.AppId, c.AppKey)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("format", Format)
	req.Header.Set("charset", Charset)

	resp, err := c.HttpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	xlog.Debug("uri:", addr, " params:", string(body), " response:", string(respBody))
	if err != nil {
		return nil, err
	}
	c.SetResponseBody(respBody)
	return respBody, nil
}

// 包装的通用请求和解析
func (c *Client) doRequest(reqObj, respObj interface{}, addr string) error {
	respBody, err := c.requestApi(reqObj, addr)
	if err != nil {
		return err
	}
	return c.decodeRespObj(respBody, respObj)
}

// 解析通用响应结构体 并且验证网关和业务状态码
func (c *Client) decodeRespObj(respBody []byte, respObj interface{}) error {
	// 解析应答结构体
	if err := json.Unmarshal(respBody, respObj); err != nil {
		return err
	}
	// 解析通用应答字段，判断网关状态码和业务状态码
	commonResp := new(CommonResponseParams)
	if err := json.Unmarshal(respBody, commonResp); err != nil {
		return err
	}

	// 判断网关状态码
	if commonResp.ErrCode != GateWaySuccess {
		msg := commonResp.ErrMsg + " " + commonResp.ErrInfo
		return c.GetGateWayError(commonResp.ErrCode, msg)
	}

	return nil
}

//AliMiniAppPreOrder 支付宝支付下单-支付宝(支付宝跳转小程序支付)
func (c *Client) AliMiniAppPreOrder(req *AliPreOrderRequest) (respObj *PreOrderResponse, err error) {
	respObj = new(PreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, AliMiniAppPreOrderApi)
	} else {
		err = c.doRequest(req, respObj, AliMiniAppPreOrderApiBeta)
	}

	return respObj, err
}

//QmfAppPreOrder 全民支付-支付下单
func (c *Client) QmfAppPreOrder(req *WxPreOrderRequest) (respObj *PreOrderResponse, err error) {
	respObj = new(PreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, QmfPreOrderApi)
	} else {
		err = c.doRequest(req, respObj, QmfPreOrderApiBeta)
	}

	return respObj, err
}

//UacAppPreOrder 银联云闪付-支付下单
func (c *Client) UacAppPreOrder(req *WxPreOrderRequest) (respObj *PreOrderResponse, err error) {
	respObj = new(PreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, UacPreOrderApi)
	} else {
		err = c.doRequest(req, respObj, UacPreOrderApiBeta)
	}

	return respObj, err
}

//AppleAppPreOrder Apple pay-支付下单
func (c *Client) AppleAppPreOrder(req *WxPreOrderRequest) (respObj *PreOrderResponse, err error) {
	respObj = new(PreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, ApplePreOrderApi)
	} else {
		err = c.doRequest(req, respObj, ApplePreOrderApiBeta)
	}

	return respObj, err
}

//WxPreOrder 微信支付下单
func (c *Client) WxPreOrder(req *WxPreOrderRequest) (respObj *PreOrderResponse, err error) {
	respObj = new(PreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, WxPreOrderApiBeta)
	} else {
		err = c.doRequest(req, respObj, WxPreOrderApi)
	}

	return respObj, err
}

//AliPreOrder 支付宝支付下单
func (c *Client) AliPreOrder(req *AliPreOrderRequest) (respObj *PreOrderResponse, err error) {
	respObj = new(PreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, AliPreOrderApiBeta)
	} else {
		err = c.doRequest(req, respObj, AliPreOrderApi)
	}

	return respObj, err
}

//ClosePreOrder 关闭支付下单
func (c *Client) ClosePreOrder(req *ClosePreOrderRequest) (respObj *ClosePreOrderResponse, err error) {
	respObj = new(ClosePreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, PreOrderCloseApi)
	} else {
		err = c.doRequest(req, respObj, PreOrderCloseApiBeta)
	}

	return respObj, err
}

//CloseWxPreOrder 关闭支付下单-微信(微信跳转小程序专用)
func (c *Client) CloseWxPreOrder(req *ClosePreOrderRequest) (respObj *ClosePreOrderResponse, err error) {
	respObj = new(ClosePreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, WxMiniAppPreOrderCloseApi)
	} else {
		err = c.doRequest(req, respObj, WxMiniAppPreOrderCloseApiBeta)
	}

	return respObj, err
}

//CloseAliPreOrder 关闭支付下单-支付宝(支付宝跳转小程序专用)
func (c *Client) CloseAliPreOrder(req *ClosePreOrderRequest) (respObj *ClosePreOrderResponse, err error) {
	respObj = new(ClosePreOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, AliMiniAppPreOrderCloseApi)
	} else {
		err = c.doRequest(req, respObj, AliMiniAppPreOrderCloseApiBeta)
	}

	return respObj, err
}

//Refund 退款
func (c *Client) Refund(req *RefundRequest) (respObj *RefundResponse, err error) {
	respObj = new(RefundResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, RefundApi)
	} else {
		err = c.doRequest(req, respObj, RefundApiBeta)
	}

	return respObj, err
}

//PreOrderQuery 交易订单查询
func (c *Client) PreOrderQuery(req *QueryOrderRequest) (respObj *QueryOrderResponse, err error) {
	respObj = new(QueryOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, OrderQueryApi)
	} else {
		err = c.doRequest(req, respObj, OrderQueryApiBeta)
	}

	return respObj, err
}

//WxPreOrderQuery 交易订单查询-微信(微信跳转小程序专用)
func (c *Client) WxPreOrderQuery(req *QueryOrderRequest) (respObj *QueryOrderResponse, err error) {
	respObj = new(QueryOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, WxMiniAppOrderQueryApi)
	} else {
		err = c.doRequest(req, respObj, WxMiniAppOrderQueryApiBeta)
	}

	return respObj, err
}

//AliPreOrderQuery 交易订单查询-支付宝(支付宝跳转小程序专用)
func (c *Client) AliPreOrderQuery(req *QueryOrderRequest) (respObj *QueryOrderResponse, err error) {
	respObj = new(QueryOrderResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, AliMiniAppOrderQueryApi)
	} else {
		err = c.doRequest(req, respObj, AliMiniAppOrderQueryApiBeta)
	}

	return respObj, err
}

//RefundQuery 退款订单查询
func (c *Client) RefundQuery(req *QueryRefundRequest) (respObj *QueryRefundResponse, err error) {
	respObj = new(QueryRefundResponse)

	if c.IsProEnv {
		err = c.doRequest(req, respObj, RefundQueryApi)
	} else {
		err = c.doRequest(req, respObj, RefundQueryApiBeta)
	}

	return respObj, err
}
