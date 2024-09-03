package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscSignQueryContractParam struct {
	SignId string `json:"signId,omitempty"`
}

func NewSmscSignQueryContractParam(signId string) *SmscSignQueryContractParam {
	return &SmscSignQueryContractParam{
		SignId: signId,
	}
}

func (c *Config) SmscSignQueryContractRequest(param *SmscSignQueryContractParam) (data xmap.M, err error) {
	method := "smsc.sign.queryContract"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
