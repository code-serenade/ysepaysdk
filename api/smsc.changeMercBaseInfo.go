package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscChangeMercBaseInfoParam struct {
	CustId            string   `json:"custId,omitempty"`            // 客户号id,custId，mercId二选一，若都传则以商户号为准
	MercId            string   `json:"mercId,omitempty"`            // 商户号,custId，mercId二选一，若都传则以商户号为准
	NotifyUrl         string   `json:"notifyUrl,omitempty"`         // 基本信息变更异步通知url,若为空，取资料上送的异步通知url
	CustInfo          CustInfo `json:"custInfo,omitempty"`          // 客户基本信息
	BusInfo           BusInfo  `json:"busInfo,omitempty"`           // 营业信息
	CrpInfo           CrpInfo  `json:"crpInfo,omitempty"`           // 法人信息
	ContractType      string   `json:"contractType,omitempty"`      // 变更申请单类型,1 纸质 2电子 默认1
	ChangeThirdFlowId string   `json:"changeThirdFlowId,omitempty"` // 业务方变更申请流水号,唯一 （变更 状态查询接口可通过此字段查询）
}

func NewSmscChangeMercBaseInfoParam() *SmscChangeMercBaseInfoParam {
	return &SmscChangeMercBaseInfoParam{}
}

func (c *Config) SmscChangeMercBaseInfoRequest(param *SmscChangeMercBaseInfoParam) (data xmap.M, err error) {
	method := "smsc.changeMercBaseInfo"
	version := "1.3"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
