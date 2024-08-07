package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type ReportScanUnionAddParam struct {
	ChannelCode             string `json:"channelCode"`             // 交易报备渠道编号
	MercId                  string `json:"mercId"`                  // 商户号
	LinencePhoto            string `json:"linencePhoto"`            // 营业执照复印件
	IndentityPhoto          string `json:"indentityPhoto"`          // 法人身份证复印件
	ProtocolPhoto           string `json:"protocolPhoto"`           // 服务协议复印件
	Cocc                    string `json:"cocc"`                    // 组织机构代码证
	SfzFrontPhoto           string `json:"sfzFrontPhoto"`           // 法人身份证正面照
	SfzBackPhoto            string `json:"sfzBackPhoto"`            // 法人身份证反面照
	IdcarNo                 string `json:"idcarNo"`                 // 法人身份证号码
	MccSubCd                string `json:"mccSubCd"`                // 银盛商户类型编码
	BusiType                string `json:"busiType"`                // 业务类型
	MercShortName           string `json:"mercShortName"`           // 商户简称
	MercProv                string `json:"mercProv"`                // 商户归属省编码
	MercCity                string `json:"mercCity"`                // 商户归属市编码
	MercArea                string `json:"mercArea"`                // 商户归属区编码
	BusAddr                 string `json:"busAddr"`                 // 营业地址
	ContactsName            string `json:"contactsName"`            // 商户联系人名称
	ContactsTel             string `json:"contactsTel"`             // 商户联系人手机
	ContactsEmail           string `json:"contactsEmail"`           // 商户联系人邮箱
	CrpNm                   string `json:"crpNm"`                   // 法人名称
	OrgNo                   string `json:"orgNo"`                   // 商户归属机构
	CertNo                  string `json:"certNo"`                  // 法人证件号码
	AgentName               string `json:"agentName"`               // 代理商名称
	AgtMercId               string `json:"agtMercId"`               // 代理商编号
	BankType                string `json:"bankType"`                // 开户行行别
	BankName                string `json:"bankName"`                // 开户行名称
	AccountType             string `json:"accountType"`             // 账户类型
	AccountName             string `json:"accountName"`             // 账户名称
	AccountNo               string `json:"accountNo"`               // 账户账号
	CustId                  string `json:"custId"`                  // 客户号
	MchType                 string `json:"mchType"`                 // 商户类型
	MercName                string `json:"mercName"`                // 商户名称
	ReportName              string `json:"reportName"`              // 微信报备名称
	MercType                string `json:"mercType"`                // 商户类别
	CertType                string `json:"certType"`                // 法人证件类型
	PId                     string `json:"pId"`                     // 支付宝推荐关系pid
	ShopPhoto               string `json:"shopPhoto"`               // 门头照
	StoreEnvirPhoto         string `json:"storeEnvirPhoto"`         // 店内环境照
	IdImgHand               string `json:"idImgHand"`               // 手持身份证照
	BankCardImg             string `json:"bankCardImg"`             // 银行卡照片
	BdShopPhoto             string `json:"bdShopPhoto"`             // 客户经理与门头照合影
	BdPosterPhoto           string `json:"bdPosterPhoto"`           // 客户经理与门店摇摇乐海报合影
	WxPlatform              string `json:"wxPlatform"`              // 平台入驻照片(微信)
	IdValidDateBegin        string `json:"idValidDateBegin"`        // 法人证件有效开始期
	IdValidDateEnd          string `json:"idValidDateEnd"`          // 法人证件有效结束期
	BusLincenceBegin        string `json:"busLincenceBegin"`        // 营业执照有效开始期
	BusLincenceEnd          string `json:"busLincenceEnd"`          // 营业执照失效期
	ManagementType          string `json:"managementType"`          // 小微商户经营类型
	StoreName               string `json:"storeName"`               // 门店名称
	WechatCheckStandPhoto   string `json:"wechatCheckStandPhoto"`   // 微信收银台照片
	AlipayCheckStandPhoto   string `json:"alipayCheckStandPhoto"`   // 支付宝收银台主扫照片
	AliPayStoreCashierPhoto string `json:"aliPayStoreCashierPhoto"` // 支付宝收银台被扫照片
	BookType                string `json:"bookType"`                // 证书类型
	UnitPhoto               string `json:"unitPhoto"`               // 单位证明函
	HotLine                 string `json:"hotLine"`                 // 客服电话
	BusinessLicenseType     string `json:"businessLicenseType"`     // 证照类型
	BusinessLicense         string `json:"businessLicense"`         // 证件编号
	AppletAppId             string `json:"appletAppId"`             // 小程序APPID
	AppId1                  string `json:"appId1"`                  // 公众号1
	AppId2                  string `json:"appId2"`                  // 公众号2
	AppId3                  string `json:"appId3"`                  // 公众号3
	AppId4                  string `json:"appId4"`                  // 公众号4
	ApplyServices           string `json:"applyServices"`           // 申请的服务
	PicType                 string `json:"picType"`                 // 上传图片类型
	JsapiPath1              string `json:"jsapiPath1"`              // 公众号授权目录1
	JsapiPath2              string `json:"jsapiPath2"`              // 公众号授权目录2
	JsapiPath3              string `json:"jsapiPath3"`              // 公众号授权目录3
	JsapiPath4              string `json:"jsapiPath4"`              // 公众号授权目录4
	JsapiPath5              string `json:"jsapiPath5"`              // 公众号授权目录5
}

// 4.2 必填选项为Y的参数
func NewReportScanUnionAddParam(channelCode, mercId, mccSubCd, busiType, mercShortName, mercProv, mercCity, mercArea, busAddr, contactsName, contactsTel, contactsEmail, crpNm, orgNo, certNo, bankType, bankName, accountType, accountName, accountNo, custId, mchType, mercType, businessLicenseType, businessLicense string) *ReportScanUnionAddParam {
	return &ReportScanUnionAddParam{
		ChannelCode:         channelCode,
		MercId:              mercId,
		MccSubCd:            mccSubCd,
		BusiType:            busiType,
		MercShortName:       mercShortName,
		MercProv:            mercProv,
		MercCity:            mercCity,
		MercArea:            mercArea,
		BusAddr:             busAddr,
		ContactsName:        contactsName,
		ContactsTel:         contactsTel,
		ContactsEmail:       contactsEmail,
		CrpNm:               crpNm,
		OrgNo:               orgNo,
		CertNo:              certNo,
		BankType:            bankType,
		BankName:            bankName,
		AccountType:         accountType,
		AccountName:         accountName,
		AccountNo:           accountNo,
		CustId:              custId,
		MchType:             mchType,
		MercType:            mercType,
		BusinessLicenseType: businessLicenseType,
		BusinessLicense:     businessLicense,
	}
}

// 根据文档生成的请求方法
func (c *Config) ReportScanUnionAddRequest(param *ReportScanUnionAddParam) (data xmap.M, err error) {
	method := "report.scan.union.add"
	version := "1.0"
	url := proUrlPrefix + reportScanUnionAddUrl
	if IsDev {
		url = devUrlPrefix + reportScanUnionAddUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
