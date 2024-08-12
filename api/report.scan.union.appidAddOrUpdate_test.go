package api

import "testing"

func TestReportScanUnionAppidAddOrUpdateRequest(t *testing.T) {
	// mercId := "826588257340024"
	// appid := "wxe8ba7f0739bd18b8"

	mercId := "826332273110012"
	appid := "wx5104d04cb34f977a"

	param := NewReportScanUnionAppidAddOrUpdateParam("CUPS_WECHAT", mercId)
	param.JsPayRelatedAppId = appid

	data, err := sdk.ReportScanUnionAppidAddOrUpdateRequest(param)
	t.Log(data, err)
}
