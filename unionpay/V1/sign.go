package unionpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func Sign(A []byte, appId, appKey string) (string, error) {
	B, err := Sha256WithLowerHexStr(A)
	if err != nil {
		return "", err
	}

	nonce := NewRandomBase64(10)
	ts := time.Now().Format("20060102150405")
	C := appId + ts + nonce + B
	D := appKey
	E := HmacSha256(C, D)
	F := base64.StdEncoding.EncodeToString(E)

	return fmt.Sprintf(`OPEN-BODY-SIG AppId="%s", Timestamp="%s", Nonce="%s", Signature="%s"`, appId, ts, nonce, F), nil
}

func Sha256WithLowerHexStr(A []byte) (string, error) {
	enc := sha256.New()
	_, err := enc.Write(A)
	if err != nil {
		return "", err
	}
	return strings.ToLower(hex.EncodeToString(enc.Sum(nil))), nil
}

func HmacSha256(C, D string) []byte {
	mac := hmac.New(sha256.New, []byte(D))
	mac.Write([]byte(C))
	return mac.Sum(nil)
}
