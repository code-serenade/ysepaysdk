package api

import "testing"

func TestPregateTradeFindCmmtAreaInfoListRequest(t *testing.T) {
	param := NewPregateTradeFindCmmtAreaInfoListParam("北京")
	param.PageNumber = "1"
	param.PageSize = "1"
	data, err := sdk.PregateTradeFindCmmtAreaInfoListRequest(param)
	t.Log(data, err)
}
