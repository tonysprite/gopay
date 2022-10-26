package unionpay

import (
	"github.com/go-pay/gopay/pkg/xlog"
	"testing"
)

func Test_GetOrderId(t *testing.T) {
	sourceCode := "A014"
	orderId := "2323343434"
	newOrderId := GetAppPreOrderId(sourceCode, orderId)
	if len(newOrderId) > 28 {
		xlog.Error("错误======== GetAppPreOrderId length over limit 28," + newOrderId)
	}
	if len(newOrderId) < 6 {
		xlog.Error("错误======== GetAppPreOrderId length less than limit 6," + newOrderId)
	}
	t.Logf("%+v", newOrderId)
}
