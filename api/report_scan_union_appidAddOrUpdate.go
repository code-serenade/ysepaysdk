package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type ReportScanUnionAppidAddOrUpdateParam struct {
	// 	channelId	String	Y	交易报备渠道编号,目前只支持微信，CUPS_WECHAT-银联微信；NUCC_WECHAT-网联微信；
	// mercId	String	Y	企业商户号
	// jsPayRelatedAppId	String	N	公众号appid,公众号appid、小程序appid二者必选一
	// appletId	String	N	小程序appid,公众号appid、小程序appid二者必选一
	ChannelId         string `json:"channelId"`
	MercId            string `json:"mercId"`
	JsPayRelatedAppId string `json:"jsPayRelatedAppId"`
	AppletId          string `json:"appletId"`
}

func NewReportScanUnionAppidAddOrUpdateParam(channelId, mercId string) *ReportScanUnionAppidAddOrUpdateParam {
	return &ReportScanUnionAppidAddOrUpdateParam{
		ChannelId: channelId,
		MercId:    mercId,
	}
}

func (c *Config) ReportScanUnionAppidAddOrUpdateRequest(param *ReportScanUnionAppidAddOrUpdateParam) (data xmap.M, err error) {
	method := "report.scan.union.appIdAddOrUpdate"
	version := "1.0"
	url := proUrlPrefix + reportScanUnionAppIdAddOrUpdateUrl
	if IsDev {
		url = devUrlPrefix + reportScanUnionAppIdAddOrUpdateUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)

	// 结构转化
	if err != nil {
		// TODO respon
		return
	}
	return
}
