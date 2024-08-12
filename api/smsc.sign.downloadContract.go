package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscSignDownloadContractParam struct {
	SignId string `json:"signId,omitempty"`
}

func NewSmscSignDownloadContractParam(signId string) *SmscSignDownloadContractParam {
	return &SmscSignDownloadContractParam{
		SignId: signId,
	}
}

func (c *Config) SmscSignDownloadContractRequest(param *SmscSignDownloadContractParam) (data xmap.M, err error) {
	method := "smsc.sign.downloadContract"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
