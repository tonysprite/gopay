## unionpay

### 银联商务支付接口
#### V1版本支持接口清单
* 下单接口-支付宝
* 下单接口-微信
* 下单接口-全民支付
* 下单接口-银联云闪付
* 下单接口-Apple Pay
* 下单接口-支付宝(支付宝跳转小程序支付)
* 订单交易查询
* 订单交易查询-微信(微信跳转小程序专用)
* 订单交易查询-支付宝(支付宝跳转小程序专用)
* 订单关闭-用户创建订单之后，对未支付的订单进行关闭操作
* 订单关闭-微信(微信跳转小程序专用)
* 订单关闭-支付宝(支付宝跳转小程序支付)
* 退款
* 退款查询

> 具体使用参考 unionpay/V1/client_test.go

> 测试的时候记得换一下这个配置 
> 
> unionpay/V1/client_test.go line:10
```go

const (
	appKey       = "xxxxxx"
	appid        = "xxxxxx"
	merchantCode = "xxxxx"
	terminalCode = "xxxxx"
	isProEnv     = false
)
```