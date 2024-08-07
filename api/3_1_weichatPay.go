package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type WeChatPayRequestParam struct {
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
	SubOpenId         string            `json:"subOpenId"`         // 用户在商户appid下的唯⼀标识
	IsMiniPg          string            `json:"isMiniPg"`          // 是否为⼩程序支付
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
	SubMerchantInfo   SubMerchantInfo   `json:"subMerchantInfo"`   // 二级商户信息
	ConsigneeInfo     ConsigneeInfo     `json:"consigneeInfo"`     // 收货人信息
	Detail            Detail            `json:"detail"`            // 微信商品详情
	UnionQrcodeParams UnionQrcodeParams `json:"unionQrcodeParams"` // 银联259号文条码改造参数
	BuyerRealNameInfo BuyerRealNameInfo `json:"buyerRealNameInfo"` // 实名认证信息
	MsgCode           string            `json:"msgCode"`           // 报文编号
}

// 4.2.1 subMerchantInfo 具体参数
type SubMerchantInfo struct {
	MerName      string `json:"merName"`      // 二级商户名称
	MerShortName string `json:"merShortName"` // 二级商户简称
	MerAddr      string `json:"merAddr"`      // 二级商户地址
	Telephone    string `json:"telephone"`    // 客户手机号
	MerNo        string `json:"merNo"`        // 二级商户编号
	Category     string `json:"category"`     // 类目编号
	MrchntCertId string `json:"mrchntCertId"` // 身份证号
}

// 4.2.2 consigneeInfo 具体参数
type ConsigneeInfo struct {
	ConsigneeName      string `json:"consigneeName"`      // 收货人姓名
	ConsigneeAddr      string `json:"consigneeAddr"`      // 收货地址
	TransportationInfo string `json:"transportationInfo"` // 物流配送信息
	CommodityName      string `json:"commodityName"`      // 商品名称
	CommodityNumber    string `json:"commodityNumber"`    // 商品数量
}

// 4.2.3 detail 具体参数
type Detail struct {
	CostPrice    string        `json:"costPrice"`    // 订单原价
	ReceiptId    string        `json:"receiptId"`    // 商品小票ID
	GoodsDetails []GoodsDetail `json:"goodsDetails"` // 商品列表
}

// 4.2.3.1 goodsDetails 具体参数
type GoodsDetail struct {
	GoodsId      string `json:"goodsId"`      // 商户商品编码
	WxpayGoodsId string `json:"wxpayGoodsId"` // 微信商品编号
	GoodsName    string `json:"goodsName"`    // 商品名称
	Quantity     string `json:"quantity"`     // 商品数量
	Price        string `json:"price"`        // 商品单价
}

// 4.2.4 unionQrcodeParams 具体参数
type UnionQrcodeParams struct {
	TerminalNo     string `json:"terminalNo"`     // 终端号
	TerminalType   string `json:"terminalType"`   // 终端类型
	SerialNum      string `json:"serialNum"`      // 终端序列号
	NetworkLicense string `json:"networkLicense"` // 入网认证编号
	AppVersion     string `json:"appVersion"`     // 终端应用版本号
	TerminalGps    string `json:"terminalGps"`    // 交易设备位置信息
	TerminalIp     string `json:"terminalIp"`     // 终端设备IP
	EncrypRandNum  string `json:"encrypRandNum"`  // 加密随机因子
	SecretText     string `json:"secretText"`     // 密文数据
}

// 4.2.5 buyerRealNameInfo 具体参数
type BuyerRealNameInfo struct {
	IdType string `json:"idType"` // 证件类型
	IdName string `json:"idName"` // 证件名称
	IdNo   string `json:"idNo"`   // 证件号码
}

// 创建 WeChatPayRequestParam 实例
func NewWeChatPayRequestParam(orderId, note, amount, payeeMercId, timeOutExpress, busiCode, subOpenId, appId string) *WeChatPayRequestParam {
	return &WeChatPayRequestParam{
		OrderId:        orderId,
		Note:           note,
		Amount:         amount,
		PayeeMercId:    payeeMercId,
		TimeOutExpress: timeOutExpress,
		BusiCode:       busiCode,
		SubOpenId:      subOpenId,
		AppId:          appId,
	}
}

// 根据文档生成的请求方法
func (c *Config) WeChatPayRequest(param *WeChatPayRequestParam) (data xmap.M, err error) {
	method := "unify.basePay.scan.weChatPay.js"
	version := "1.0"
	url := proUrlPrefix + weChatPayUrl
	if IsDev {
		url = devUrlPrefix + weChatPayUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
