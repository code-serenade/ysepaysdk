package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type T1SmscAddCustInfoApplyParam struct {
	CustInfo    CustInfo   `json:"custInfo,omitempty"`    // 基本信息
	CrpInfo     CrpInfo    `json:"crpInfo,omitempty"`     // 法人信息
	StlAccInfo  StlAccInfo `json:"stlAccInfo,omitempty"`  // 结算信息
	BusInfo     BusInfo    `json:"busInfo,omitempty"`     // 营业信息
	ThirdFlowId string     `json:"thirdFlowId,omitempty"` // 业务方入网申请流水号
}

// 4.2.1 custInfo 具体参数
type CustInfo struct {
	AgtMercId     string `json:"agtMercId,omitempty"`     // 代理商编号
	MercName      string `json:"mercName,omitempty"`      // 商户名称
	MercShortName string `json:"mercShortName,omitempty"` // 商户简称
	MercType      string `json:"mercType,omitempty"`      // 商户类型,2 小微 3个体 4企业 5 社会组织 6 事业单位 7 政府机关
	MccCd         string `json:"mccCd,omitempty"`         // mcc码
	ContactMail   string `json:"contactMail,omitempty"`   // 联系人邮箱
	ContactMan    string `json:"contactMan,omitempty"`    // 联系人
	ContactPhone  string `json:"contactPhone,omitempty"`  // 联系人电话
	CusMgrNm      string `json:"cusMgrNm,omitempty"`      // 客户经理
	IsOpenMarket  string `json:"isOpenMarket,omitempty"`  // 是否开通营销
	NotifyUrl     string `json:"notifyUrl,omitempty"`     // 异步通知地址
	Remark        string `json:"remark,omitempty"`        // 备注
}

// 4.2.2 crpInfo 具体参数
type CrpInfo struct {
	CrpCertNo    string       `json:"crpCertNo,omitempty"`    // 法人证件号
	CrpCertType  string       `json:"crpCertType,omitempty"`  // 法人证件类型
	CertBgn      string       `json:"certBgn,omitempty"`      // 证件开始日期
	CertExpire   string       `json:"certExpire,omitempty"`   // 证件有效期
	CrpNati      string       `json:"crpNati,omitempty"`      // 法人国籍
	CrpNm        string       `json:"crpNm,omitempty"`        // 法人姓名
	CrpPhone     string       `json:"crpPhone,omitempty"`     // 法人手机号
	CrpAddr      string       `json:"crpAddr,omitempty"`      // 法人地址
	ActContrInfo ActContrInfo `json:"actContrInfo,omitempty"` // 实际控制人
	AuthInfo     AuthInfo     `json:"authInfo,omitempty"`     // 被授权人信息
	BnfList      []BnfInfo    `json:"bnfList,omitempty"`      // 受益人列表
}

// 4.2.2.1 实际控制人信息
type ActContrInfo struct {
	ContrNm         string `json:"contrNm,omitempty"`         // 实际控制人姓名
	ContrCertType   string `json:"contrCertType,omitempty"`   // 实际控制人证件类型
	ContrCertNo     string `json:"contrCertNo,omitempty"`     // 实际控制人证件号码
	ContrCertBgn    string `json:"contrCertBgn,omitempty"`    // 实际控制人证件开始日期
	ContrCertExpire string `json:"contrCertExpire,omitempty"` // 实际控制人证件有效期
}

// 4.2.2.2 被授权人信息
type AuthInfo struct {
	AuthNm         string `json:"authNm,omitempty"`         // 被授权人姓名
	AuthCertType   string `json:"authCertType,omitempty"`   // 被授权人证件类型
	AuthCertNo     string `json:"authCertNo,omitempty"`     // 被授权人证件号码
	AuthCertBgn    string `json:"authCertBgn,omitempty"`    // 被授权人证件开始日期
	AuthCertExpire string `json:"authCertExpire,omitempty"` // 被授权人证件有效期
}

// 4.2.2.3 受益人信息
type BnfInfo struct {
	BnfNm         string `json:"bnfNm,omitempty"`         // 受益人姓名
	BnfCertType   string `json:"bnfCertType,omitempty"`   // 受益人证件类型
	BnfCertNo     string `json:"bnfCertNo,omitempty"`     // 受益人证件号码
	BnfCertBgn    string `json:"bnfCertBgn,omitempty"`    // 受益人证件开始日期
	BnfCertExpire string `json:"bnfCertExpire,omitempty"` // 受益人证件有效期
}

// 4.2.3 stlAccInfo 具体参数
type StlAccInfo struct {
	IsSettInPlatAcc     string `json:"isSettInPlatAcc,omitempty"`     // 是否平台内账户
	IsUncrpSett         string `json:"isUncrpSett,omitempty"`         // 是否非法人结算
	StlAccNm            string `json:"stlAccNm,omitempty"`            // 结算户名
	StlAccNo            string `json:"stlAccNo,omitempty"`            // 结算账号
	BankSubCode         string `json:"bankSubCode,omitempty"`         // 开户支行联行号
	StlAccType          string `json:"stlAccType,omitempty"`          // 结算账户类型
	BankMobile          string `json:"bankMobile,omitempty"`          // 银行预留手机号
	OpenCertNo          string `json:"openCertNo,omitempty"`          // 开开户证件号
	BankProince         string `json:"bankProince,omitempty"`         // 银行开户行所属省代码
	BankCity            string `json:"bankCity,omitempty"`            // 银行开户行所属市代码
	StandByStlAccType   string `json:"standByStlAccType,omitempty"`   // 备用结算账户类型
	StandByStlAccNo     string `json:"standByStlAccNo,omitempty"`     // 备用结算账号
	StandByStlAccNm     string `json:"standByStlAccNm,omitempty"`     // 备用结算户名
	StandByBankSubCode  string `json:"standByBankSubCode,omitempty"`  // 备用结算账户开户行支行编号
	StandByBankProvince string `json:"standByBankProvince,omitempty"` // 备用结算账户银行开户行省代码
	StandByBankCity     string `json:"standByBankCity,omitempty"`     // 备用结算账户银行开户行市代码
}

// 4.2.4 busInfo 具体参数
type BusInfo struct {
	BusNo          string `json:"busNo,omitempty"`          // 营业执照号
	BusNm          string `json:"busNm,omitempty"`          // 营业执照名称
	BusCertBgn     string `json:"busCertBgn,omitempty"`     // 营业执照有效开始日期
	BusCertExpire  string `json:"busCertExpire,omitempty"`  // 营业执照有效期
	BusProviceCode string `json:"busProviceCode,omitempty"` // 营业归属省代码
	BusCityCode    string `json:"busCityCode,omitempty"`    // 营业归属市代码
	BusAreaCode    string `json:"busAreaCode,omitempty"`    // 营业归属区(县)代码
	BusCertType    string `json:"busCertType,omitempty"`    // 营业执照证件类型
	BusAddr        string `json:"busAddr,omitempty"`        // 营业详细地址
	RegAddr        string `json:"regAddr,omitempty"`        // 注册地址
}

// 创建 T1SmscAddCustInfoApplyParam 实例
func NewT1SmscAddCustInfoApplyParam(custInfo CustInfo, crpInfo CrpInfo, stlAccInfo StlAccInfo, busInfo BusInfo, thirdFlowId string) *T1SmscAddCustInfoApplyParam {
	return &T1SmscAddCustInfoApplyParam{
		CustInfo:    custInfo,
		CrpInfo:     crpInfo,
		StlAccInfo:  stlAccInfo,
		BusInfo:     busInfo,
		ThirdFlowId: thirdFlowId,
	}
}

// 根据文档生成的请求方法
func (c *Config) T1SmscAddCustInfoApplyRequest(param *T1SmscAddCustInfoApplyParam) (data xmap.M, err error) {
	method := "t1.smsc.addCustInfoApply"
	version := "1.2"
	url := proUrlPrefix + t1SmscAddCustInfoApplyUrl
	if IsDev {
		url = devUrlPrefix + t1SmscAddCustInfoApplyUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
