package ysepaysdk

import (
	"fmt"
	"testing"
	"time"
)

func TestOnlineQrcodepayRequest(t *testing.T) {
	outTradeNo := time.Now().Format("20060102150405")
	subject := "test"
	sellerID := "sellerID"
	businessCode := "00510030"
	bankType := "1903000"
	totalAmount := "1.00"
	partnerID := "partnerID"
	notifyURL := "https://example.com/v1/ysepay/notifyPay"
	bizContent := NewBizContentOnlineQrcodepay(outTradeNo, subject, sellerID, businessCode, bankType, totalAmount)
	resp, err := Shared.OnlineQrcodepayRequest(partnerID, notifyURL, bizContent)
	fmt.Printf("%#v %v", resp, err)
}
