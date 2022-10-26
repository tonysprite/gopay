package unionpay

import (
	"github.com/go-pay/gopay/pkg/xlog"
	"testing"
	"time"
)

const (
	appKey       = "xxxxxx"
	appid        = "xxxxxx"
	merchantCode = "xxxxx"
	terminalCode = "xxxxx"
	isProEnv     = false
)

var (
	requestId = NewRandomBase64(10)
	client    = NewClient(requestId, appid, appKey, isProEnv)
)

//测试支付宝支付
func TestClient_AliPreOrder(t *testing.T) {
	req := &AliPreOrderRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       getOrderId(),
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
		TotalAmount:      1,
		TradeType:        "APP",
	}

	response, err := client.AliPreOrder(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}

	t.Logf("%+v", response)
}

//测试微信支付
func TestClient_WxPreOrder(t *testing.T) {
	req := &WxPreOrderRequest{
		MsgId:            requestId,
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       getOrderId(),
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
		TotalAmount:      1,
		TradeType:        "APP",
		SubAppId:         "wx232sf2323",
	}

	response, err := client.WxPreOrder(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

//测试关闭订单
func TestClient_ClosePreOrder(t *testing.T) {
	req := &ClosePreOrderRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666706976",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
	}

	response, err := client.ClosePreOrder(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

func TestClient_CloseWxPreOrder(t *testing.T) {
	req := &ClosePreOrderRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666706976",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
	}

	response, err := client.CloseWxPreOrder(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

func TestClient_CloseAliPreOrder(t *testing.T) {
	req := &ClosePreOrderRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666750033",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
	}

	response, err := client.CloseAliPreOrder(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

func TestClient_Refund(t *testing.T) {
	req := &RefundRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666750033",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
		RefundAmount:     1,
	}

	response, err := client.Refund(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

func TestClient_PreOrderQuery(t *testing.T) {
	req := &QueryOrderRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666750033",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
		TargetOrderId:    "103A1666750033",
	}

	response, err := client.PreOrderQuery(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

func TestClient_WxPreOrderQuery(t *testing.T) {
	req := &QueryOrderRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666750033",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
		TargetOrderId:    "103A1666750033",
	}

	response, err := client.WxPreOrderQuery(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}

func TestClient_RefundQuery(t *testing.T) {
	req := &QueryRefundRequest{
		Mid:              merchantCode,
		RequestTimestamp: time.Now().Format("2006-01-02 15:04:05"),
		MerOrderId:       "103A1666750033",
		Tid:              terminalCode,
		InstMid:          "APPDEFAULT",
	}

	response, err := client.RefundQuery(req)
	if err != nil {
		xlog.Error("错误========", err)
		return
	}
	t.Logf("%+v", response)
}
