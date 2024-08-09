package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscqueryCustApplyParam struct {
	// sysFlowId	String(25)	N	入网申请流水号
	// thirdFlowId	String(32)	N	业务方入网申请流水号,（若资料上送接口 有上送此字段，则可通过此字段查询商户 信息）
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
	url := proUrlPrefix + smscQueryCustApplyUrl
	if IsDev {
		url = devUrlPrefix + smscQueryCustApplyUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
