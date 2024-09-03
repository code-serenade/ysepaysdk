package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type CodeScanT1Fee struct {
	WxPayFee          FeeParam `json:"wxPayFee,omitempty"`          // 微信扫码费率
	AliPayFee         FeeParam `json:"aliPayFee,omitempty"`         // 支付宝扫码费率
	AliPayDebitFee    FeeParam `json:"aliPayDebitFee,omitempty"`    // 支付宝扫码借记卡费率
	AliPayCreditFee   FeeParam `json:"aliPayCreditFee,omitempty"`   // 支付宝扫码贷记卡费率
	Bank1debitPayFee  FeeParam `json:"bank1debitPayFee,omitempty"`  // 银联一档借记卡扫码费率(>1000)
	Bank1creditPayFee FeeParam `json:"bank1creditPayFee,omitempty"` // 银联一档贷记卡扫码费率(>1000)
	Bank2debitPayFee  FeeParam `json:"bank2debitPayFee,omitempty"`  // 银联二档借记卡扫码费率(<=1000)
	Bank2creditPayFee FeeParam `json:"bank2creditPayFee,omitempty"` // 银联二档贷记卡扫码费率(<=1000)
}

func NewCodeScanT1Fee() CodeScanT1Fee {
	return CodeScanT1Fee{}
}

type SwCardT1Fee struct {
	DebitPayFee  FeeParam `json:"debitPayFee,omitempty"`  // 借记卡费率
	CreditPayFee FeeParam `json:"creditPayFee,omitempty"` // 贷记卡费率
}

func NewSwCardT1Fee() SwCardT1Fee {
	return SwCardT1Fee{}
}

type SwCardD0Fee struct {
	DebitPayFee  FeeParam `json:"debitPayFee,omitempty"`  // 借记卡费率
	CreditPayFee FeeParam `json:"creditPayFee,omitempty"` // 贷记卡费率
}

func NewSwCardD0Fee() SwCardD0Fee {
	return SwCardD0Fee{}
}

type FeeParam struct {
	RateType   string `json:"rateType,omitempty"`   // 收费方式,0 按百分比 1按固定金额
	RateFee    string `json:"rateFee,omitempty"`    // 费率,按百分比时单位为% 按固定金额时 单位为分
	RateBottom string `json:"rateBottom,omitempty"` // 最低收费,单位为分 按百分比时生效
	RateTop    string `json:"rateTop,omitempty"`    // 最高收费,单位为分
}

func NewFeeParam() FeeParam {
	return FeeParam{}
}

type SmscChangeRateParam struct {
	CodeScanT1Fee     CodeScanT1Fee `json:"codeScanT1Fee,omitempty"`     // 扫码工作日到账费率,busOpenType=00时必填
	CodeScanD0Fee     FeeParam      `json:"codeScanD0Fee,omitempty"`     // 扫码实时到账垫资费,busOpenType=01时必填
	SwCardT1Fee       SwCardT1Fee   `json:"swCardT1Fee,omitempty"`       // 刷卡工作日到账费率,busOpenType=10时必填
	SwCardD0Fee       SwCardD0Fee   `json:"swCardD0Fee,omitempty"`       // 刷卡实时到账垫资费,busOpenType=11时必填
	D1Fee             FeeParam      `json:"d1Fee,omitempty"`             // 天天到账垫资费,busOpenType=20时必填
	CustId            string        `json:"custId,omitempty"`            // 客户号,custId，mercId二选一，若都传则以商户号为准
	MercId            string        `json:"mercId,omitempty"`            // 商户号,custId，mercId二选一，若都传则以商户号为准
	NotifyUrl         string        `json:"notifyUrl,omitempty"`         // 费率信息变更异步通知url,若为空，取资料上送的异步通知url
	ContractType      string        `json:"contractType,omitempty"`      // 变更申请单类型,1 纸质 2电子 默认1
	ChangeThirdFlowId string        `json:"changeThirdFlowId,omitempty"` // 业务方变更申请流水号,唯一 （变更 状态查询接口可通过此字段查询）

}

func NewSmscChangeRateParam(custId string) *SmscChangeRateParam {
	return &SmscChangeRateParam{
		CustId: custId,
	}
}

// 根据文档生成的请求方法
func (c *Config) SmscChangeRateRequest(param *SmscChangeRateParam) (data xmap.M, err error) {
	method := "smsc.changeRate"
	version := "1.4"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
