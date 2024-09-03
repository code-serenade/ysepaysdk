package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// sysFlowId	String(25)	Y	入网申请流水
// custInfo	Object	Y	商户基本信息
// crpInfo	Object	Y	法人信息
// stlAccInfo	Object	Y	结算信息
// busInfo	Object	Y	营业信息

type SmscModifyCustInfoApplyParam struct {
	SysFlowId  string     `json:"sysFlowId,omitempty"`
	CustInfo   CustInfo   `json:"custInfo,omitempty"`
	CrpInfo    CrpInfo    `json:"crpInfo,omitempty"`
	StlAccInfo StlAccInfo `json:"stlAccInfo,omitempty"`
	BusInfo    BusInfo    `json:"busInfo,omitempty"`
}

func NewSmscModifyCustInfoApplyParam() *SmscModifyCustInfoApplyParam {
	return &SmscModifyCustInfoApplyParam{}
}

func (c *Config) SmscModifyCustInfoApplyRequest(param *SmscModifyCustInfoApplyParam) (data xmap.M, err error) {
	method := "smsc.modifyCustInfoApply"
	version := "1.2"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
