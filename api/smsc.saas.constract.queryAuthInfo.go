package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscSaasConstractQueryAuthInfoParam struct {
	AuthId      string `json:"authId,omitempty"`
	ThirdAuthId string `json:"thirdAuthId,omitempty"`
}

func NewSmscSaasConstractQueryAuthInfoParam(authId, thirdAuthId string) *SmscSaasConstractQueryAuthInfoParam {
	param := &SmscSaasConstractQueryAuthInfoParam{}
	if len(authId) > 0 {
		param.AuthId = authId
	}
	if len(thirdAuthId) > 0 {
		param.ThirdAuthId = thirdAuthId
	}
	return param
}

func (c *Config) SmscSaasConstractQueryAuthInfoRequest(param *SmscSaasConstractQueryAuthInfoParam) (data xmap.M, err error) {
	method := "smsc.saas.constract.queryAuthInfo"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
