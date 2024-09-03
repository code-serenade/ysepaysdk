package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscqueryCustApplyParam struct {
	SysFlowId   string `json:"sysFlowId,omitempty"`
	ThirdFlowId string `json:"thirdFlowId,omitempty"`
}

func NewSmscqueryCustApplyParam(sysFlowId, thirdFlowId string) *SmscqueryCustApplyParam {
	param := &SmscqueryCustApplyParam{}
	if len(sysFlowId) > 0 {
		param.SysFlowId = sysFlowId
	}
	if len(thirdFlowId) > 0 {
		param.ThirdFlowId = thirdFlowId
	}
	return param
}

// 根据文档生成的请求方法
func (c *Config) SmscqueryCustApplyRequest(param *SmscqueryCustApplyParam) (data xmap.M, err error) {
	method := "smsc.queryCustApply"
	version := "1.1"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
