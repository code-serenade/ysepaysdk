package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type PregateTradeFindBankNameAndBankCodeParam struct {
	OpenBankName string `json:"openBankName"` //银行名称
	CityCode     string `json:"cityCode"`     //地区编码
}

func NewPregateTradeFindBankNameAndBankCodeParam(openBankName string, cityCode string) *PregateTradeFindBankNameAndBankCodeParam {
	return &PregateTradeFindBankNameAndBankCodeParam{
		OpenBankName: openBankName,
		CityCode:     cityCode,
	}
}

func (c *Config) PregateTradeFindBankNameAndBankCodeRequest(param *PregateTradeFindBankNameAndBankCodeParam) (data xmap.M, err error) {
	method := "pregate.trade.findBankNameAndBankCode"
	version := "1.0"
	url := proUrlPrefix + pregateTradeFindBankNameAndBankCodeUrl
	if IsDev {
		url = devUrlPrefix + pregateTradeFindBankNameAndBankCodeUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
