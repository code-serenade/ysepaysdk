package api

const (
	ALLCHAR = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	devUrlPrefix = "https://appdev.ysepay.com"
	proUrlPrefix = "https://ysgate.ysepay.com"

	fileUploadUrl = "/openapi/file/smsc/upload"

	successCode = "00000"
)

var (
	IsDev   = false
	Verbose = false
)
