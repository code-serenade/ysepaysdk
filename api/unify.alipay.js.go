package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type AlipayLifeRequestParam struct {
	OrderId           string            `json:"orderId,omitempty"`           // 商户订单号
	ShopDate          string            `json:"shopDate,omitempty"`          // 交易日期
	Note              string            `json:"note,omitempty"`              // 订单备注
	Amount            string            `json:"amount,omitempty"`            // 订单金额
	Currency          string            `json:"currency,omitempty"`          // 币种
	PayeeMercId       string            `json:"payeeMercId,omitempty"`       // 收款方商户号
	PayeeMercName     string            `json:"payeeMercName,omitempty"`     // 收款方商户名称
	TimeOutExpress    string            `json:"timeOutExpress,omitempty"`    // 订单失效时间
	ExtraCommonParam  string            `json:"extraCommonParam,omitempty"`  // 公用回传参数
	BusiCode          string            `json:"busiCode,omitempty"`          // 业务代码
	SubOpenId         string            `json:"subOpenId,omitempty"`         // 用户在商户appid下的唯一标识
	IsMiniPg          string            `json:"isMiniPg,omitempty"`          // 是否为小程序支付
	AppId             string            `json:"appId,omitempty"`             // 商户公众平台ID
	LimitCreditPay    string            `json:"limitCreditPay,omitempty"`    // 是否限制信用卡
	AllowRepeatPay    string            `json:"allowRepeatPay,omitempty"`    // 是否允许多次支付
	FailNotifyUrl     string            `json:"failNotifyUrl,omitempty"`     // 失败通知地址
	NotifyUrl         string            `json:"notifyUrl,omitempty"`         // 通知地址
	SrcIP             string            `json:"srcIP,omitempty"`             // 发起方IP
	PayerIP           string            `json:"payerIP,omitempty"`           // 付款方IP
	GpsInfo           string            `json:"gpsInfo,omitempty"`           // 付款方gps信息
	AgentNo           string            `json:"agentNo,omitempty"`           // 代理商编号
	PaygateNo         string            `json:"paygateNo,omitempty"`         // 支付网关编号
	MsgCode           string            `json:"msgCode,omitempty"`           // 报文编号
	SubMerchantInfo   SubMerchantInfo   `json:"subMerchantInfo,omitempty"`   // 二级商户信息
	ConsigneeInfo     ConsigneeInfo     `json:"consigneeInfo,omitempty"`     // 收货人信息
	Detail            Detail            `json:"detail,omitempty"`            // 商品详情
	UnionQrcodeParams UnionQrcodeParams `json:"unionQrcodeParams,omitempty"` // 银联259号文条码改造参数
	BuyerRealNameInfo BuyerRealNameInfo `json:"buyerRealNameInfo,omitempty"` // 实名认证信息
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
	method := "unify.alipay.js"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
