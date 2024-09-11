package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// mercId	String	Y	商户号,由银盛支付生成并下发
// channelId	String	Y	交易报备渠道编号,CUPS_WECHAT-银联微信；NUCC_WECHAT-网联微信；
// payAuthPath	String	Y	微信公众号添加配置支付授权目录

type ReportScanUnionAuthPathAddParam struct {
	MercId      string `json:"mercId"`
	ChannelId   string `json:"channelId"`
	PayAuthPath string `json:"payAuthPath"`
}

func NewReportScanUnionAuthPathAddParam(mercId, channelId, payAuthPath string) *ReportScanUnionAuthPathAddParam {
	return &ReportScanUnionAuthPathAddParam{
		MercId:      mercId,
		ChannelId:   channelId,
		PayAuthPath: payAuthPath,
	}
}

// 这段代码需要帮我生成
func (c *Config) ReportScanUnionAuthPathAddRequest(param *ReportScanUnionAuthPathAddParam) (data xmap.M, err error) {
	method := "report.scan.union.authPathAdd"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
