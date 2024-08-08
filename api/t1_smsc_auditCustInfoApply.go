package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type T1SmscAuditCustInfoApplyParam struct {
	AuditFlag string `json:"auditFlag,omitempty"` // 审核标志：Y-通过，N-拒绝
	SysFlowId string `json:"sysFlowId,omitempty"` // 入网申请流水号
}

func NewT1SmscAuditCustInfoApplyParam(sysFlowId string) *T1SmscAuditCustInfoApplyParam {
	return &T1SmscAuditCustInfoApplyParam{
		AuditFlag: "Y",
		SysFlowId: sysFlowId,
	}
}

func (c *Config) T1SmscAuditCustInfoApplyRequest(param *T1SmscAuditCustInfoApplyParam) (data xmap.M, err error) {
	method := "t1.smsc.auditCustInfoApply"
	version := "1.0"
	url := proUrlPrefix + t1SmscAuditCustInfoApplyUrl
	if IsDev {
		url = devUrlPrefix + t1SmscAuditCustInfoApplyUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
