package api

const (
	ALLCHAR = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// api 接口地址
	devUrlPrefix                       = "https://appdev.ysepay.com"
	proUrlPrefix                       = "https://ysgate.ysepay.com"
	fileUploadUrl                      = "/openapi/file/smsc/upload"
	t1SmscAddCustInfoApplyUrl          = "/openapi/t1/smsc/addCustInfoApply"
	reportScanUnionAppIdAddOrUpdateUrl = "/openapi/report/scan/union/appIdAddOrUpdate"

	successCode = "00000"
)

var (
	IsDev   = false
	Verbose = false
)
