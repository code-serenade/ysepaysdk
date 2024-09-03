package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type ReportScanUnionAppidAddOrUpdateParam struct {
	ChannelId         string `json:"channelId,omitempty"`         // 交易报备渠道编号,目前只支持微信，CUPS_WECHAT-银联微信；NUCC_WECHAT-网联微信；
	MercId            string `json:"mercId,omitempty"`            // 企业商户号
	JsPayRelatedAppId string `json:"jsPayRelatedAppId,omitempty"` // 公众号appid,公众号appid、小程序appid二者必选一
	AppletId          string `json:"appletId,omitempty"`          // 小程序appid,公众号appid、小程序appid二者必选一
}

// 4.2 必填选项为Y的参数
func NewReportScanUnionAppidAddOrUpdateParam(channelId, mercId string) *ReportScanUnionAppidAddOrUpdateParam {
	return &ReportScanUnionAppidAddOrUpdateParam{
		ChannelId: channelId,
		MercId:    mercId,
	}
}

// 这段代码需要帮我生成
func (c *Config) ReportScanUnionAppidAddOrUpdateRequest(param *ReportScanUnionAppidAddOrUpdateParam) (data xmap.M, err error) {
	method := "report.scan.union.appIdAddOrUpdate"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
