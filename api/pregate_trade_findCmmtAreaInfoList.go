package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type PregateTradeFindCmmtAreaInfoListParam struct {
	PageNumber string `json:"pageNumber,omitempty"` // 当前页码
	PageSize   string `json:"pageSize,omitempty"`   // 每页显示条数,最大50
	CityName   string `json:"cityName,omitempty"`   // 地区码名称
	CityCd     string `json:"cityCd,omitempty"`     // 地区码
}

func NewPregateTradeFindCmmtAreaInfoListParam(cityName string) *PregateTradeFindCmmtAreaInfoListParam {
	return &PregateTradeFindCmmtAreaInfoListParam{
		PageNumber: "1",
		PageSize:   "50",
		CityName:   cityName,
	}
}

func (c *Config) PregateTradeFindCmmtAreaInfoListRequest(param *PregateTradeFindCmmtAreaInfoListParam) (data xmap.M, err error) {
	method := "pregate.trade.findCmmtAreaInfoList"
	version := "1.0"
	url := proUrlPrefix + pregateTradeFindCmmtAreaInfoListUrl
	if IsDev {
		url = devUrlPrefix + pregateTradeFindCmmtAreaInfoListUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
