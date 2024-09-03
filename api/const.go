package api

import "strings"

const (
	ALLCHAR = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// api 接口地址
	devUrlPrefix   = "https://appdev.ysepay.com"
	proUrlPrefix   = "https://ysgate.ysepay.com"
	successCode    = "00000"
	subSuccessCode = "0000"
)

var (
	IsDev   = false // 是否开发环境
	Verbose = false // for debug
)

// return host/openapi/path
func methodToUrl(method string) string {
	path := methodToPath(method)
	if IsDev {
		return devUrlPrefix + "/openapi/" + path
	}
	return proUrlPrefix + "/openapi/" + path
}

func methodToPath(method string) string {
	return strings.ReplaceAll(method, ".", "/")
}
