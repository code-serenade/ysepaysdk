package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type WeChatPayRequestParam struct {
	OrderId           string             `json:"orderId,omitempty"`           // 商户订单号
	ShopDate          string             `json:"shopDate,omitempty"`          // 交易日期
	Note              string             `json:"note,omitempty"`              // 订单备注
	Amount            string             `json:"amount,omitempty"`            // 订单金额
	Currency          string             `json:"currency,omitempty"`          // 币种
	PayeeMercId       string             `json:"payeeMercId,omitempty"`       // 收款方商户号
	PayeeMercName     string             `json:"payeeMercName,omitempty"`     // 收款方商户名称
	TimeOutExpress    string             `json:"timeOutExpress,omitempty"`    // 订单失效时间
	ExtraCommonParam  string             `json:"extraCommonParam,omitempty"`  // 公用回传参数
	BusiCode          string             `json:"busiCode,omitempty"`          // 业务代码
	SubOpenId         string             `json:"subOpenId,omitempty"`         // 用户在商户appid下的唯⼀标识
	IsMiniPg          string             `json:"isMiniPg,omitempty"`          // 是否为⼩程序支付
	AppId             string             `json:"appId,omitempty"`             // 商户公众平台ID
	LimitCreditPay    string             `json:"limitCreditPay,omitempty"`    // 是否限制信用卡
	AllowRepeatPay    string             `json:"allowRepeatPay,omitempty"`    // 是否允许多次支付
	FailNotifyUrl     string             `json:"failNotifyUrl,omitempty"`     // 失败通知地址
	NotifyUrl         string             `json:"notifyUrl,omitempty"`         // 通知地址
	SrcIP             string             `json:"srcIP,omitempty"`             // 发起方IP
	PayerIP           string             `json:"payerIP,omitempty"`           // 付款方IP
	GpsInfo           string             `json:"gpsInfo,omitempty"`           // 付款方gps信息
	AgentNo           string             `json:"agentNo,omitempty"`           // 代理商编号
	PaygateNo         string             `json:"paygateNo,omitempty"`         // 支付网关编号
	SubMerchantInfo   *SubMerchantInfo   `json:"subMerchantInfo,omitempty"`   // 二级商户信息
	ConsigneeInfo     *ConsigneeInfo     `json:"consigneeInfo,omitempty"`     // 收货人信息
	Detail            *Detail            `json:"detail,omitempty"`            // 微信商品详情
	UnionQrcodeParams *UnionQrcodeParams `json:"unionQrcodeParams,omitempty"` // 银联259号文条码改造参数
	BuyerRealNameInfo *BuyerRealNameInfo `json:"buyerRealNameInfo,omitempty"` // 实名认证信息
	MsgCode           string             `json:"msgCode,omitempty"`           // 报文编号
}

// 4.2.1 subMerchantInfo 具体参数
type SubMerchantInfo struct {
	MerName      string `json:"merName,omitempty"`      // 二级商户名称
	MerShortName string `json:"merShortName,omitempty"` // 二级商户简称
	MerAddr      string `json:"merAddr,omitempty"`      // 二级商户地址
	Telephone    string `json:"telephone,omitempty"`    // 客户手机号
	MerNo        string `json:"merNo,omitempty"`        // 二级商户编号
	Category     string `json:"category,omitempty"`     // 类目编号
	MrchntCertId string `json:"mrchntCertId,omitempty"` // 身份证号
}

// 4.2.2 consigneeInfo 具体参数
type ConsigneeInfo struct {
	ConsigneeName      string `json:"consigneeName,omitempty"`      // 收货人姓名
	ConsigneeAddr      string `json:"consigneeAddr,omitempty"`      // 收货地址
	TransportationInfo string `json:"transportationInfo,omitempty"` // 物流配送信息
	CommodityName      string `json:"commodityName,omitempty"`      // 商品名称
	CommodityNumber    string `json:"commodityNumber,omitempty"`    // 商品数量
}

// 4.2.3 detail 具体参数
type Detail struct {
	CostPrice    string        `json:"costPrice,omitempty"`    // 订单原价
	ReceiptId    string        `json:"receiptId,omitempty"`    // 商品小票ID
	GoodsDetails []GoodsDetail `json:"goodsDetails,omitempty"` // 商品列表
}

// 4.2.3.1 goodsDetails 具体参数
type GoodsDetail struct {
	GoodsId      string `json:"goodsId,omitempty"`      // 商户商品编码
	WxpayGoodsId string `json:"wxpayGoodsId,omitempty"` // 微信商品编号
	GoodsName    string `json:"goodsName,omitempty"`    // 商品名称
	Quantity     string `json:"quantity,omitempty"`     // 商品数量
	Price        string `json:"price,omitempty"`        // 商品单价
}

// 4.2.4 unionQrcodeParams 具体参数
type UnionQrcodeParams struct {
	TerminalNo     string `json:"terminalNo,omitempty"`     // 终端号
	TerminalType   string `json:"terminalType,omitempty"`   // 终端类型
	SerialNum      string `json:"serialNum,omitempty"`      // 终端序列号
	NetworkLicense string `json:"networkLicense,omitempty"` // 入网认证编号
	AppVersion     string `json:"appVersion,omitempty"`     // 终端应用版本号
	TerminalGps    string `json:"terminalGps,omitempty"`    // 交易设备位置信息
	TerminalIp     string `json:"terminalIp,omitempty"`     // 终端设备IP
	EncrypRandNum  string `json:"encrypRandNum,omitempty"`  // 加密随机因子
	SecretText     string `json:"secretText,omitempty"`     // 密文数据
}

// 4.2.5 buyerRealNameInfo 具体参数
type BuyerRealNameInfo struct {
	IdType string `json:"idType,omitempty"` // 证件类型
	IdName string `json:"idName,omitempty"` // 证件名称
	IdNo   string `json:"idNo,omitempty"`   // 证件号码
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
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
