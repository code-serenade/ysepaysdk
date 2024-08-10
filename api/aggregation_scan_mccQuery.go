package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type AggregationScanMccQueryParam struct {
	MccCd    string `json:"mccCd,omitempty"`    // mcc码
	MercType string `json:"mercType,omitempty"` // 商户类型,2 小微 3个体 4企业 5 社会组织 6 事业单位 7 政府机关
}

func NewAggregationScanMccQueryParam(mercType string) *AggregationScanMccQueryParam {
	return &AggregationScanMccQueryParam{
		MercType: mercType,
	}
}

func (c *Config) AggregationScanMccQueryRequest(param *AggregationScanMccQueryParam) (data xmap.M, err error) {
	method := "aggregation.scan.mccQuery"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
