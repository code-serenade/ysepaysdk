package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// signId	String(20)	Y	签约流水号
// isSendConMsg	String(1)	Y	通知方式,0 短信+邮件 1 短信 2邮件 3不通知

type SmscSignSendSmsOrEmailMsgParam struct {
	SignId       string `json:"signId,omitempty"`
	IsSendConMsg string `json:"isSendConMsg,omitempty"`
}

func NewSmscSignSendSmsOrEmailMsgParam(signId, isSendConMsg string) *SmscSignSendSmsOrEmailMsgParam {
	return &SmscSignSendSmsOrEmailMsgParam{
		SignId:       signId,
		IsSendConMsg: isSendConMsg,
	}
}

func (c *Config) SmscSignSendSmsOrEmailMsgRequest(param *SmscSignSendSmsOrEmailMsgParam) (data xmap.M, err error) {
	method := "smsc.sign.sendSmsOrEmailMsg"
	version := "1.0"
	url := methodToUrl(method)

	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
