package api

import (
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// 4.2 业务请求参数
type ReportScanUnionUploadParam struct {
	CustInfo    CustInfo   `json:"custInfo"`    // 基本信息
	CrpInfo     CrpInfo    `json:"crpInfo"`     // 法人信息
	StlAccInfo  StlAccInfo `json:"stlAccInfo"`  // 结算信息
	BusInfo     BusInfo    `json:"busInfo"`     // 营业信息
	ThirdFlowId string     `json:"thirdFlowId"` // 业务方入网申请流水号
}

// 4.2.1 custInfo 具体参数
type CustInfo struct {
	AgtMercId     string `json:"agtMercId"`     // 代理商编号
	MercName      string `json:"mercName"`      // 商户名称
	MercShortName string `json:"mercShortName"` // 商户简称
	MercType      string `json:"mercType"`      // 商户类型
	MccCd         string `json:"mccCd"`         // mcc码
	ContactMail   string `json:"contactMail"`   // 联系人邮箱
	ContactMan    string `json:"contactMan"`    // 联系人
	ContactPhone  string `json:"contactPhone"`  // 联系人电话
	CusMgrNm      string `json:"cusMgrNm"`      // 客户经理
	IsOpenMarket  string `json:"isOpenMarket"`  // 是否开通营销
	NotifyUrl     string `json:"notifyUrl"`     // 异步通知地址
	Remark        string `json:"remark"`        // 备注
}

// 4.2.2 crpInfo 具体参数
type CrpInfo struct {
	CrpCertNo    string       `json:"crpCertNo"`    // 法人证件号
	CrpCertType  string       `json:"crpCertType"`  // 法人证件类型
	CertBgn      string       `json:"certBgn"`      // 证件开始日期
	CertExpire   string       `json:"certExpire"`   // 证件有效期
	CrpNati      string       `json:"crpNati"`      // 法人国籍
	CrpNm        string       `json:"crpNm"`        // 法人姓名
	CrpPhone     string       `json:"crpPhone"`     // 法人手机号
	CrpAddr      string       `json:"crpAddr"`      // 法人地址
	ActContrInfo ActContrInfo `json:"actContrInfo"` // 实际控制人
	AuthInfo     AuthInfo     `json:"authInfo"`     // 被授权人信息
	BnfList      []BnfInfo    `json:"bnfList"`      // 受益人列表
}

// 4.2.2.1 实际控制人信息
type ActContrInfo struct {
	ContrNm         string `json:"contrNm"`         // 实际控制人姓名
	ContrCertType   string `json:"contrCertType"`   // 实际控制人证件类型
	ContrCertNo     string `json:"contrCertNo"`     // 实际控制人证件号码
	ContrCertBgn    string `json:"contrCertBgn"`    // 实际控制人证件开始日期
	ContrCertExpire string `json:"contrCertExpire"` // 实际控制人证件有效期
}

// 4.2.2.2 被授权人信息
type AuthInfo struct {
	AuthNm         string `json:"authNm"`         // 被授权人姓名
	AuthCertType   string `json:"authCertType"`   // 被授权人证件类型
	AuthCertNo     string `json:"authCertNo"`     // 被授权人证件号码
	AuthCertBgn    string `json:"authCertBgn"`    // 被授权人证件开始日期
	AuthCertExpire string `json:"authCertExpire"` // 被授权人证件有效期
}

// 4.2.2.3 受益人信息
type BnfInfo struct {
	BnfNm         string `json:"bnfNm"`         // 受益人姓名
	BnfCertType   string `json:"bnfCertType"`   // 受益人证件类型
	BnfCertNo     string `json:"bnfCertNo"`     // 受益人证件号码
	BnfCertBgn    string `json:"bnfCertBgn"`    // 受益人证件开始日期
	BnfCertExpire string `json:"bnfCertExpire"` // 受益人证件有效期
}

// 4.2.3 stlAccInfo 具体参数
type StlAccInfo struct {
	IsSettInPlatAcc     string `json:"isSettInPlatAcc"`     // 是否平台内账户
	IsUncrpSett         string `json:"isUncrpSett"`         // 是否非法人结算
	StlAccNm            string `json:"stlAccNm"`            // 结算户名
	StlAccNo            string `json:"stlAccNo"`            // 结算账号
	BankSubCode         string `json:"bankSubCode"`         // 开户支行联行号
	StlAccType          string `json:"stlAccType"`          // 结算账户类型
	BankMobile          string `json:"bankMobile"`          // 银行预留手机号
	OpenCertNo          string `json:"openCertNo"`          // 开开户证件号
	BankProince         string `json:"bankProince"`         // 银行开户行所属省代码
	BankCity            string `json:"bankCity"`            // 银行开户行所属市代码
	StandByStlAccType   string `json:"standByStlAccType"`   // 备用结算账户类型
	StandByStlAccNo     string `json:"standByStlAccNo"`     // 备用结算账号
	StandByStlAccNm     string `json:"standByStlAccNm"`     // 备用结算户名
	StandByBankSubCode  string `json:"standByBankSubCode"`  // 备用结算账户开户行支行编号
	StandByBankProvince string `json:"standByBankProvince"` // 备用结算账户银行开户行省代码
	StandByBankCity     string `json:"standByBankCity"`     // 备用结算账户银行开户行市代码
}

// 4.2.4 busInfo 具体参数
type BusInfo struct {
	BusNo          string `json:"busNo"`          // 营业执照号
	BusNm          string `json:"busNm"`          // 营业执照名称
	BusCertBgn     string `json:"busCertBgn"`     // 营业执照有效开始日期
	BusCertExpire  string `json:"busCertExpire"`  // 营业执照有效期
	BusProviceCode string `json:"busProviceCode"` // 营业归属省代码
	BusCityCode    string `json:"busCityCode"`    // 营业归属市代码
	BusAreaCode    string `json:"busAreaCode"`    // 营业归属区(县)代码
	BusCertType    string `json:"busCertType"`    // 营业执照证件类型
	BusAddr        string `json:"busAddr"`        // 营业详细地址
	RegAddr        string `json:"regAddr"`        // 注册地址
}

// 创建 ReportScanUnionUploadParam 实例
func NewReportScanUnionUploadParam(custInfo CustInfo, crpInfo CrpInfo, stlAccInfo StlAccInfo, busInfo BusInfo, thirdFlowId string) *ReportScanUnionUploadParam {
	return &ReportScanUnionUploadParam{
		CustInfo:    custInfo,
		CrpInfo:     crpInfo,
		StlAccInfo:  stlAccInfo,
		BusInfo:     busInfo,
		ThirdFlowId: thirdFlowId,
	}
}

// 根据文档生成的请求方法
func (c *Config) ReportScanUnionUploadRequest(param *ReportScanUnionUploadParam) (data xmap.M, err error) {
	method := "t1.smsc.addCustInfoApply"
	version := "1.2"
	url := proUrlPrefix + reportScanUnionUploadUrl
	if IsDev {
		url = devUrlPrefix + reportScanUnionUploadUrl
	}
	bizContent := converter.JSON(param)
	_, data, err = c.Request(url, method, version, bizContent)
	return
}
