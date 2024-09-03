package ysepaysdk

import (
	"github.com/codingeasygo/util/xmap"
)

func (c Config) OnlineQrcodepayRequest(partnerID, notifyURL string, bizContent *BizContentOnlineQrcodepay) (resp xmap.M, err error) {
	req := NewAPIRequest("ysepay.online.qrcodepay", partnerID, notifyURL, bizContent.JSONString())
	return c.Request(QrcodeAPIUrl, req)
}
