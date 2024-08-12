package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type T1SmscSignParam struct {
	BusOpenType   string        `json:"busOpenType,omitempty"`   // 到账方式  00-扫码工作日到账01-扫码实时到账10-刷卡工作日到账11-刷卡实时到账20-D1到账 允许多选用“|”分隔 T1工作日到账必选
	ContractType  string        `json:"contractType,omitempty"`  // 合同类型  1-纸质合同、2-电子合同，其他证明文件、基金会法人登记证书仅支持纸质合同，其他均支持电子合同、纸质合同。
	IsSendConMsg  string        `json:"isSendConMsg,omitempty"`  // 是否自动发送签约通知,contractType=2时必填 0(短信+邮件) 1(短信) 2(邮件) 3(不通知)
	NotifyUrl     string        `json:"notifyUrl,omitempty"`     // 异步通知url
	CustId        string        `json:"custId,omitempty"`        // 客户号
	CodeScanD0Fee FeeParam      `json:"codeScanD0Fee,omitempty"` // 扫码实时到账垫资费,busOpenType=01时必填
	CodeScanT1Fee CodeScanT1Fee `json:"codeScanT1Fee,omitempty"` // 扫码工作日到账费率,busOpenType=00时必填
	SwCardD0Fee   SwCardT1Fee   `json:"swCardD0Fee,omitempty"`   // 刷卡实时到账垫资费,busOpenType=11时必填
	SwCardT1Fee   SwCardT1Fee   `json:"swCardT1Fee,omitempty"`   // 刷卡工作日到账费率,busOpenType=10时必填
	D1Fee         FeeParam      `json:"d1Fee,omitempty"`         // 天天到账垫资费,busOpenType=20时必填
	ThirdAuthId   string        `json:"thirdAuthId,omitempty"`   // 业务方权限id,唯一（签约状态信息接口 可通过此字段查询）
}

func NewT1SmscSignParam() *T1SmscSignParam {
	return &T1SmscSignParam{}
}

func (c *Config) T1SmscSignRequest(param *T1SmscSignParam) (data xmap.M, err error) {
	method := "t1.smsc.sign"
	version := "1.2"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
