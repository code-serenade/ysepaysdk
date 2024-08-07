package api

import "testing"

func TestPregateTradeFindCmmtAreaInfoListRequest(t *testing.T) {

	param := NewPregateTradeFindCmmtAreaInfoListParam("上海市")

	data, err := sdk.PregateTradeFindCmmtAreaInfoListRequest(param)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

}
