package api

import (
	"time"

	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type UnifyScanQrcodePayRequest struct {
	OrderId           string             `json:"orderId,omitempty"`
	ShopDate          string             `json:"shopDate,omitempty"`
	Note              string             `json:"note,omitempty"`
	Amount            string             `json:"amount,omitempty"`
	Currency          string             `json:"currency,omitempty"`
	PayeeMercId       string             `json:"payeeMercId,omitempty"`
	PayeeMercName     string             `json:"payeeMercName,omitempty"`
	NotifyUrl         string             `json:"notifyUrl,omitempty"`
	ReturnUrl         string             `json:"returnUrl,omitempty"`
	TimeOutExpress    string             `json:"timeOutExpress,omitempty"`
	ExtendParams      string             `json:"extendParams,omitempty"`
	ExtraCommonParam  string             `json:"extraCommonParam,omitempty"`
	BusiCode          string             `json:"busiCode,omitempty"`
	BankType          string             `json:"bankType,omitempty"`
	SubMerchantInfo   *SubMerchantInfo   `json:"subMerchantInfo,omitempty"`
	ConsigneeInfo     *ConsigneeInfo     `json:"consigneeInfo,omitempty"`
	LimitCreditPay    string             `json:"limitCreditPay,omitempty"`
	HbFqNum           string             `json:"hbFqNum,omitempty"`
	AllowRepeatPay    string             `json:"allowRepeatPay,omitempty"`
	FqType            string             `json:"fqType,omitempty"`
	AliGoodsDetails   []AliGoodsDetail   `json:"aliGoodsDetails,omitempty"`
	SubmerIp          string             `json:"submerIp,omitempty"`
	StoreId           string             `json:"storeId,omitempty"`
	AlipayStoreId     string             `json:"alipayStoreId,omitempty"`
	OperatorId        string             `json:"operatorId,omitempty"`
	TerminalId        string             `json:"terminalId,omitempty"`
	UnionQrcodeParams *UnionQrcodeParams `json:"unionQrcodeParams,omitempty"`
	NeedHttpsUrl      string             `json:"needHttpsUrl,omitempty"`
	BuyerRealnameInfo *BuyerRealnameInfo `json:"buyerRealnameInfo,omitempty"`
	MsgCode           string             `json:"msgCode,omitempty"`
}

type AliGoodsDetail struct {
	GoodsId        string `json:"goodsId,omitempty"`
	AlipayGoodsId  string `json:"alipayGoodsId,omitempty"`
	GoodsName      string `json:"goodsName,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	Price          string `json:"price,omitempty"`
	GoodsCategory  string `json:"goodsCategory,omitempty"`
	CategoriesTree string `json:"categoriesTree,omitempty"`
	Body           string `json:"body,omitempty"`
	ShowUrl        string `json:"showUrl,omitempty"`
}

type BuyerRealnameInfo struct {
	IdNo   string `json:"idNo,omitempty"`
	IdType string `json:"idType,omitempty"`
	IdName string `json:"idName,omitempty"`
}

func NewUnifyScanQrcodePayRequest(orderId, note, amount, payeeMercId, notifyUrl, timeOutExpress, busiCode, bankType string) *UnifyScanQrcodePayRequest {
	return &UnifyScanQrcodePayRequest{
		OrderId:        orderId,
		ShopDate:       time.Now().Format("20060102"),
		Note:           note,
		Amount:         amount,
		Currency:       "CNY",
		PayeeMercId:    payeeMercId,
		NotifyUrl:      notifyUrl,
		TimeOutExpress: timeOutExpress,
		BusiCode:       busiCode,
		BankType:       bankType,
	}
}

func (c *Config) UnifyScanQrcodePay(param *UnifyScanQrcodePayRequest) (data xmap.M, err error) {
	method := "unify.scan.qrcodepay"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
