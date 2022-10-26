package unionpay

import (
	"math/rand"
	"strconv"
	"time"
)

//const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//NewRandomBase64 dataLen: 数据长度，base64膨胀率1.33,返回的string长度是dataLen的1.33倍
func NewRandomBase64(strLen int) string {
	ba := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		ba[i] = encodeURL[rand.Int()%len(encodeURL)]
	}

	return string(ba)
}

//func NewRandomHex(strLen int) string {
//	ba := make([]byte, strLen)
//	baseStr := []byte("0123456789abcdef")
//	for i := 0; i < strLen; i++ {
//		ba[i] = baseStr[rand.Int()%len(baseStr)]
//	}
//
//	return string(ba)
//}

func getOrderId() string {
	prefix := "103A"
	orderId := time.Now().Unix()
	return prefix + strconv.Itoa(int(orderId))
}

func GetAppPreOrderId(sourceCode string, orderSn string) (appOrderId string) {
	appOrderId = sourceCode + orderSn + strconv.Itoa(rand.Intn(100))
	if len(appOrderId) < 6 {
		orderId := time.Now().Unix()
		appOrderId = appOrderId + strconv.Itoa(int(orderId)) + strconv.Itoa(rand.Intn(100))
	}
	if len(appOrderId) > 28 {
		orderId := time.Now().Unix()
		appOrderId = appOrderId + strconv.Itoa(int(orderId)) + strconv.Itoa(rand.Intn(100))
	}
	return appOrderId
}
