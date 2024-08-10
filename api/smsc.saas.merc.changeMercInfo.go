package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type BaseInfoAO struct {
	CustInfo CustInfo `json:"custInfo,omitempty"`
	BusInfo  BusInfo  `json:"busInfo,omitempty"`
	CrpInfo  CrpInfo  `json:"crpInfo,omitempty"`
}

func NewBaseInfoAO() BaseInfoAO {
	return BaseInfoAO{}
}

type FeeInfoAO struct {
	CodeScanT1Fee CodeScanT1Fee `json:"codeScanT1Fee,omitempty"`
	CodeScanD0Fee FeeParam      `json:"codeScanD0Fee,omitempty"`
	SwCardT1Fee   SwCardT1Fee   `json:"swCardT1Fee,omitempty"`
	SwCardD0Fee   SwCardD0Fee   `json:"swCardD0Fee,omitempty"`
	D1Fee         FeeParam      `json:"d1Fee,omitempty"`
}

func NewFeeInfoAO() FeeInfoAO {
	return FeeInfoAO{}
}

type SmscSaasMercChangeMercInfoParam struct {
	CustId            string     `json:"custId,omitempty"`
	MercId            string     `json:"mercId,omitempty"`
	ContractType      string     `json:"contractType,omitempty"`
	ChangeThirdFlowId string     `json:"changeThirdFlowId,omitempty"`
	NotifyUrl         string     `json:"notifyUrl,omitempty"`
	BaseInfoAO        BaseInfoAO `json:"baseInfoAO,omitempty"`
	StlAccInfo        StlAccInfo `json:"stlAccInfo,omitempty"`
	FeeInfoAO         FeeInfoAO  `json:"feeInfoAO,omitempty"`
}

func NewSmscSaasMercChangeMercInfoParam() *SmscSaasMercChangeMercInfoParam {
	return &SmscSaasMercChangeMercInfoParam{}
}

func (c *Config) SmscSaasMercChangeMercInfoRequest(param *SmscSaasMercChangeMercInfoParam) (data xmap.M, err error) {
	method := "smsc.saas.merc.changeMercInfo"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
