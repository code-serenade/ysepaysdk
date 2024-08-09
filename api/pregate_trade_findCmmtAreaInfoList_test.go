package api

import "testing"

func TestPregateTradeFindCmmtAreaInfoListRequest(t *testing.T) {
	param := &PregateTradeFindCmmtAreaInfoListParam{
		PageNumber: "2",
		PageSize:   "1",
		CityCd:     "5800",
	}
	data, err := sdk.PregateTradeFindCmmtAreaInfoListRequest(param)
	t.Log(data, err)
}
