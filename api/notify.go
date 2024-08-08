package api

import (
	"bytes"
	"fmt"
	"log"

	"github.com/code-serenade/easycrypto"
	"github.com/code-serenade/ysepaysdk/utils"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/web"
)

var NotifyCustDone = func(orderID string, result xmap.M) (err error) {
	panic("NotifyCustDone not implemented")
}

var NotifyPayDone = func(orderID string, result xmap.M) (err error) {
	panic("NotifyPayDone not implemented")
}

var FindPubKey = func(orderID string) (key []byte, err error) {
	err = fmt.Errorf("FindPubKey not implemented")
	return
}

type NotifyPayload struct {
	TimeStamp  string `json:"timeStamp,omitempty"`
	Src        string `json:"src,omitempty"`
	ReqID      string `json:"reqId,omitempty"`
	Charset    string `json:"charset,omitempty"`
	Sign       string `json:"sign,omitempty"`
	BizContent string `json:"bizContent,omitempty"`
}

func (p *NotifyPayload) VerifySign(key []byte) (err error) {
	m := xmap.M{
		"timeStamp":  p.TimeStamp,
		"src":        p.Src,
		"reqId":      p.ReqID,
		"charset":    p.Charset,
		"bizContent": p.BizContent,
	}
	content := utils.MapToUrlValues(m)
	return easycrypto.RSAVerify(key, []byte(content), p.Sign)
}

// Handle 函数的逻辑是处理和分发 HTTP 请求到相应的处理程序。
// 它使用了一个正则表达式模式来匹配 URL，并将匹配到的请求分发到对应的处理程序。
func Handle(pre string, mux *web.SessionMux) {
	mux.HandleFunc("^"+pre+"/v1/ysepay/notifyCust(\\?.*)?$", handleNotify(NotifyCustDone))
	mux.HandleFunc("^"+pre+"/v1/ysepay/notifyPay(\\?.*)?$", handleNotify(NotifyPayDone))
}

func handleNotify(notifyFunc func(orderID string, result xmap.M) error) web.HandlerFunc {
	return func(s *web.Session) web.Result {
		payload, err := receiveAndVerifyPayload(s)
		if err != nil {
			log.Printf("error receiving and verifying payload: %v", err)
			return s.SendPlainText("fail")
		}

		result, err := parseBizContent(payload.BizContent)
		if err != nil {
			log.Printf("error parsing bizContent: %v", err)
			return s.SendPlainText("success")
		}

		if notifyFunc != nil {
			if err := notifyFunc(payload.ReqID, result); err != nil {
				log.Printf("notification function error: %v", err)
				return s.SendPlainText("fail")
			}
		}

		return s.SendPlainText("success")
	}
}

func receiveAndVerifyPayload(s *web.Session) (*NotifyPayload, error) {
	payload := &NotifyPayload{}
	data, err := s.RecvJSON(payload)
	if err != nil {
		log.Printf("recv json %v err:%v", string(data), err)
		return nil, err
	}

	key, err := FindPubKey(payload.ReqID)
	if err == nil {
		if err := payload.VerifySign(key); err != nil {
			log.Printf("verify sign %v err:%v", payload.ReqID, err)
			return nil, err
		}
	}
	return payload, nil
}

func parseBizContent(bizContent string) (xmap.M, error) {
	result := xmap.M{}
	if _, err := converter.UnmarshalJSON(bytes.NewBuffer([]byte(bizContent)), &result); err != nil {
		return nil, err
	}
	return result, nil
}
