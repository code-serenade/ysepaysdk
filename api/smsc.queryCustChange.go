package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscQueryCustChangeParam struct {
	ChangeSysFlowId   string `json:"changeSysFlowId,omitempty"`   // 变更申请流水
	ChangeThirdFlowId string `json:"changeThirdFlowId,omitempty"` // 业务方变更申请流水号,（若信息变更申请接口有上送此字段，则 可通过此字段查询商户变更申请信 息）
}

func NewSmscQueryCustChangeParam() *SmscQueryCustChangeParam {
	return &SmscQueryCustChangeParam{}
}

func (c *Config) SmscQueryCustChangeRequest(param *SmscQueryCustChangeParam) (data xmap.M, err error) {
	method := "smsc.queryCustChange"
	version := "1.2"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
