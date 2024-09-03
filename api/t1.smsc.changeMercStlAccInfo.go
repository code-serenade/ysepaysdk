package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type T1SmscChangeMercStlAccInfoParam struct {
	CustId            string     `json:"custId,omitempty"`            // 客户ID,custId，mercId二选一，若都传则以商户号为准
	MercId            string     `json:"mercId,omitempty"`            // 商户号,custId，mercId二选一，若都传则以商户号为准
	NotifyUrl         string     `json:"notifyUrl,omitempty"`         // 结算信息变更异步通知url,若为空，取资料上送的异步通知url
	StlAccInfo        StlAccInfo `json:"stlAccInfo,omitempty"`        // 结算信息变更
	ContractType      string     `json:"contractType,omitempty"`      // 变更申请单类型,1 纸质 2电子 默认1
	ChangeThirdFlowId string     `json:"changeThirdFlowId,omitempty"` // 业务方变更申请流水号,唯一 （变更 状态查询接口可通过此字段查询）
}

func NewT1SmscChangeMercStlAccInfoParam(custId string, stlAccInfo StlAccInfo) *T1SmscChangeMercStlAccInfoParam {
	return &T1SmscChangeMercStlAccInfoParam{
		CustId:     custId,
		StlAccInfo: stlAccInfo,
	}
}

func (c *Config) T1SmscChangeMercStlAccInfoRequest(param *T1SmscChangeMercStlAccInfoParam) (data xmap.M, err error) {
	method := "t1.smsc.changeMercStlAccInfo"
	version := "1.3"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
