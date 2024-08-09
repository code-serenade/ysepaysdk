package api

const (
	ALLCHAR = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// api 接口地址

	devUrlPrefix                       = "https://appdev.ysepay.com"
	proUrlPrefix                       = "https://ysgate.ysepay.com"
	fileUploadUrl                      = "/openapi/file/smsc/upload"
	t1SmscAddCustInfoApplyUrl          = "/openapi/t1/smsc/addCustInfoApply"
	reportScanUnionAppIdAddOrUpdateUrl = "/openapi/report/scan/union/appIdAddOrUpdate"
	reportScanUnionAddUrl              = "/openapi/report/scan/union/reportAdd"

	t1SmscAuditCustInfoApplyUrl            = "/openapi/t1/smsc/auditCustInfoApply"
	pregateTradeFindCmmtAreaInfoListUrl    = "/openapi/pregate/trade/findCmmtAreaInfoList"
	pregateTradeFindBankNameAndBankCodeUrl = "/openapi/pregate/trade/findBankNameAndBankCode"
	t1SmscSignUrl                          = "/openapi/t1/smsc/sign"

	weChatPayUrl  = "/openapi/unify/basePay/scan/weChatPay/js" // 新增的地址
	alipayLifeUrl = "/openapi/unify/alipay/js"

	aggregationScanMccQueryUrl = "/openapi/aggregation/scan/mccQuery"
	smscQueryCustApplyUrl      = "/openapi/smsc/queryCustApply"

	successCode    = "00000"
	subSuccessCode = "0000"
)

var (
	IsDev   = false // 是否开发环境
	Verbose = false // for debug
)
