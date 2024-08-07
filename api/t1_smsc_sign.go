package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type T1SmscSignParam struct {
	BusOpenType   string `json:"busOpenType"`   // 到账方式  00-扫码工作日到账01-扫码实时到账10-刷卡工作日到账11-刷卡实时到账20-D1到账 允许多选用“|”分隔 T1工作日到账必选
	ContractType  string `json:"contractType"`  // 合同类型  1-纸质合同、2-电子合同，其他证明文件、基金会法人登记证书仅支持纸质合同，其他均支持电子合同、纸质合同。
	IsSendConMsg  string `json:"isSendConMsg"`  // 是否自动发送签约通知,contractType=2时必填 0(短信+邮件) 1(短信) 2(邮件) 3(不通知)
	NotifyUrl     string `json:"notifyUrl"`     // 异步通知url
	CustId        string `json:"custId"`        // 客户号
	CodeScanD0Fee string `json:"codeScanD0Fee"` // 扫码实时到账垫资费,busOpenType=01时必填
	// TODO 这里只添加了扫码实时到账参数
}

type CodeScanD0FeeParam struct {
	RateType   string `json:"rateType"`   // 收费方式,只支持0 按百分比
	RateFee    string `json:"rateFee"`    // 费率,单位为% 填0.53代表0.53%
	RateBottom string `json:"rateBottom"` // 最低收费,单位为：分
}

func (c *Config) T1SmscSignRequest(param *T1SmscSignParam) (data xmap.M, err error) {
	method := "t1.smsc.sign"
	version := "1.2"
	url := proUrlPrefix + t1SmscSignUrl
	if IsDev {
		url = devUrlPrefix + t1SmscSignUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
