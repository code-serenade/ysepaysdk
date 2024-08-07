package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type AlipayLifeRequestParam struct {
	OrderId           string            `json:"orderId"`           // 商户订单号
	ShopDate          string            `json:"shopDate"`          // 交易日期
	Note              string            `json:"note"`              // 订单备注
	Amount            string            `json:"amount"`            // 订单金额
	Currency          string            `json:"currency"`          // 币种
	PayeeMercId       string            `json:"payeeMercId"`       // 收款方商户号
	PayeeMercName     string            `json:"payeeMercName"`     // 收款方商户名称
	TimeOutExpress    string            `json:"timeOutExpress"`    // 订单失效时间
	ExtraCommonParam  string            `json:"extraCommonParam"`  // 公用回传参数
	BusiCode          string            `json:"busiCode"`          // 业务代码
	SubOpenId         string            `json:"subOpenId"`         // 用户在商户appid下的唯一标识
	IsMiniPg          string            `json:"isMiniPg"`          // 是否为小程序支付
	AppId             string            `json:"appId"`             // 商户公众平台ID
	LimitCreditPay    string            `json:"limitCreditPay"`    // 是否限制信用卡
	AllowRepeatPay    string            `json:"allowRepeatPay"`    // 是否允许多次支付
	FailNotifyUrl     string            `json:"failNotifyUrl"`     // 失败通知地址
	NotifyUrl         string            `json:"notifyUrl"`         // 通知地址
	SrcIP             string            `json:"srcIP"`             // 发起方IP
	PayerIP           string            `json:"payerIP"`           // 付款方IP
	GpsInfo           string            `json:"gpsInfo"`           // 付款方gps信息
	AgentNo           string            `json:"agentNo"`           // 代理商编号
	PaygateNo         string            `json:"paygateNo"`         // 支付网关编号
	MsgCode           string            `json:"msgCode"`           // 报文编号
	SubMerchantInfo   SubMerchantInfo   `json:"subMerchantInfo"`   // 二级商户信息
	ConsigneeInfo     ConsigneeInfo     `json:"consigneeInfo"`     // 收货人信息
	Detail            Detail            `json:"detail"`            // 商品详情
	UnionQrcodeParams UnionQrcodeParams `json:"unionQrcodeParams"` // 银联259号文条码改造参数
	BuyerRealNameInfo BuyerRealNameInfo `json:"buyerRealNameInfo"` // 实名认证信息
}

// 创建 AlipayLifeRequestParam 实例
func NewAlipayLifeRequestParam(orderId, note, amount, payeeMercId, timeOutExpress, busiCode, notifyUrl string) *AlipayLifeRequestParam {
	return &AlipayLifeRequestParam{
		OrderId:        orderId,
		Note:           note,
		Amount:         amount,
		PayeeMercId:    payeeMercId,
		TimeOutExpress: timeOutExpress,
		BusiCode:       busiCode,
		NotifyUrl:      notifyUrl,
	}
}

// 根据文档生成的请求方法
func (c *Config) AlipayLifeRequest(param *AlipayLifeRequestParam) (data xmap.M, err error) {
	method := "unify.basePay.scan.alipay.life"
	version := "1.0"
	url := proUrlPrefix + alipayLifeUrl
	if IsDev {
		url = devUrlPrefix + alipayLifeUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
