package ysepaysdk

import (
	"bytes"
	"time"

	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 二维码支付
type BizContentOnlineQrcodepay struct {
	OutTradeNo           string                  `json:"out_trade_no"`                      // 商户系统生成的订单号
	Shopdate             string                  `json:"shopdate"`                          // 交易日期
	Subject              string                  `json:"subject"`                           // 订单备注
	TotalAmount          string                  `json:"total_amount"`                      // 资金总额
	Currency             string                  `json:"currency,omitempty"`                // 货币类型，默认CNY
	SellerID             string                  `json:"seller_id"`                         // 收款商户号
	SellerName           string                  `json:"seller_name,omitempty"`             // 收款商户名称
	TimeoutExpress       string                  `json:"timeout_express"`                   // 交易超时时间
	BusinessCode         string                  `json:"business_code"`                     // 业务代码
	BankType             string                  `json:"bank_type"`                         // 二维码行别
	ExtendParams         string                  `json:"extend_params,omitempty"`           // 业务扩展参数
	ExtraCommonParam     string                  `json:"extra_common_param,omitempty"`      // 公用回传参数
	SubMerchant          *SubMerchantInfo        `json:"sub_merchant,omitempty"`            // 二级商户信息
	ConsigneeInfo        *ConsigeeInfo           `json:"consignee_info,omitempty"`          // 收货人信息
	LimitCreditPay       string                  `json:"limit_credit_pay,omitempty"`        // 信用卡限制
	HbFqNum              string                  `json:"hb_fq_num,omitempty"`               // 花呗分期期数
	FqType               string                  `json:"fq_type,omitempty"`                 // 分期类型
	AliGoodsDetails      []AliGoodsDetail        `json:"aliGoodsDetails,omitempty"`         // 支付宝营销单品详情列表
	SubmerIP             string                  `json:"submer_ip,omitempty"`               // 子商户IP
	StoreID              string                  `json:"store_id,omitempty"`                // 商户门店编号
	AlipayStoreID        string                  `json:"alipay_store_id,omitempty"`         // 支付宝店铺编号
	OperatorID           string                  `json:"operator_id,omitempty"`             // 操作员编号
	TerminalID           string                  `json:"terminal_id,omitempty"`             // 终端编号
	UnionQrcode259Params *UnionQrCodeBaseRequest `json:"union_qrcode_259_params,omitempty"` // 银联259号文条码改造字段
	NeedHttpsURL         string                  `json:"need_https_url,omitempty"`          // 是否需要https协议图片url
	BuyerRealnameInfo    *BuyerRealnameInfo      `json:"buyer_realname_info,omitempty"`     // 实名认证信息
	ExtendParamsChannel  string                  `json:"extend_params_channel,omitempty"`   // 支付宝业务拓展参数
	CorpSrcSuperieor     string                  `json:"corp_src_superieor,omitempty"`      // 服务商信息编号
}

func NewBizContentOnlineQrcodepay(outTradeNo, subject, sellerID, businessCode, bankType, totalAmount string) *BizContentOnlineQrcodepay {
	return &BizContentOnlineQrcodepay{
		OutTradeNo:     outTradeNo,
		Shopdate:       time.Now().Format(`20060102`),
		Subject:        subject,
		TotalAmount:    totalAmount,
		SellerID:       sellerID,
		TimeoutExpress: "30m", // 交易超时时间
		BusinessCode:   businessCode,
		BankType:       bankType,
		Currency:       "CNY",
	}
}

func (b *BizContentOnlineQrcodepay) JSONString() string {
	return converter.JSON(b)
}

func (b *BizContentOnlineQrcodepay) Map() xmap.M {
	m := xmap.M{}
	converter.UnmarshalJSON(bytes.NewBufferString(converter.JSON(b)), &m)
	return m
}

func (b *BizContentOnlineQrcodepay) MapVal() (m xmap.M, err error) {
	m = xmap.M{}
	_, err = converter.UnmarshalJSON(bytes.NewBufferString(converter.JSON(b)), &m)
	return
}

type SubMerchantInfo struct {
	MerName      string `json:"merName,omitempty"`      // 二级商户名称
	MerShortName string `json:"merShortName,omitempty"` // 二级商户简称
	MerAddr      string `json:"merAddr,omitempty"`      // 二级商户地址
	Telephone    string `json:"telephone,omitempty"`    // 联系电话
	MerNo        string `json:"merNo,omitempty"`        // 二级商户编号
	Category     string `json:"category,omitempty"`     // 类目编号
	MrchntCertId string `json:"mrchntCertId,omitempty"` // 身份证号
}

type ConsigeeInfo struct {
	ConsigneeName      string `json:"consigneeName,omitempty"`      // 收货人姓名
	ConsigneeAddr      string `json:"consigneeAddr,omitempty"`      // 收货地址
	TransportationInfo string `json:"transportationInfo,omitempty"` // 物流信息
	CommodityName      string `json:"commodityName,omitempty"`      // 商品名称
	CommodityNumber    string `json:"commodityNumber,omitempty"`    // 商品数量
}

type AliGoodsDetail struct {
	GoodsID        string  `json:"goods_id"`                  // 营销单品编号
	AlipayGoodsID  string  `json:"alipay_goods_id,omitempty"` // 支付宝侧编号
	GoodsName      string  `json:"goods_name"`                // 单品名称
	Quantity       int     `json:"quantity"`                  // 数量
	Price          float64 `json:"price"`                     // 单价
	GoodsCategory  string  `json:"goods_category,omitempty"`  // 类目
	CategoriesTree string  `json:"categories_tree,omitempty"` // 类目树
	Body           string  `json:"body,omitempty"`            // 描述
	ShowURL        string  `json:"show_url,omitempty"`        // 展示地址
}

type UnionQrCodeBaseRequest struct {
	TerminalNo     string `json:"terminal_no"`               // 终端号
	TerminalType   string `json:"terminal_type"`             // 终端类型
	SerialNum      string `json:"serial_num,omitempty"`      // 序列号
	NetworkLicense string `json:"network_license,omitempty"` // 网络许可证
	AppVersion     string `json:"app_version,omitempty"`     // 应用版本
	TerminalGPS    string `json:"terminal_gps,omitempty"`    // GPS位置
	TerminalIP     string `json:"terminal_ip,omitempty"`     // 终端IP
	EncrypRandNum  string `json:"encryp_rand_num,omitempty"` // 随机因子
	SecretText     string `json:"secret_text,omitempty"`     // 密文数据
}

type BuyerRealnameInfo struct {
	IdNo   string `json:"id_no"`   // 证件号
	IdType string `json:"id_type"` // 证件类型
	IdName string `json:"id_name"` // 证件姓名
}
