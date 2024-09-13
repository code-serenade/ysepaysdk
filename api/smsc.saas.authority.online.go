package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type SmscSaasAuthorityOnlineParam struct {
	MercId     string `json:"mercId"`     // 商户号
	Option     string `json:"option"`     // 选项,ON(开通);OFF(关闭)
	RateFee    string `json:"rateFee"`    // 费率,单位为% 合同签约申请时没有开通扫码实时到账业务且没有上送扫码实时到账费率时必填
	RateBottom string `json:"rateBottom"` // 最低收费,单位为分 合同签约申请时没有开通扫码实时到账业务且没有上送扫码实时到账费率时必填
}

func NewSmscSaasAuthorityOnlineParam(mercId string) *SmscSaasAuthorityOnlineParam {
	return &SmscSaasAuthorityOnlineParam{MercId: mercId}
}

func (c *Config) SmscSaasAuthorityOnlineRequest(param *SmscSaasAuthorityOnlineParam) (data xmap.M, err error) {
	method := "smsc.saas.authority.online"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
