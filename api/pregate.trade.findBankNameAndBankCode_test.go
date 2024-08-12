package api

import "testing"

func TestPregateTradeFindBankNameAndBankCodeRequest(t *testing.T) {
	param := NewPregateTradeFindBankNameAndBankCodeParam("支行名称", "银盛地区编码")
	data, err := sdk.PregateTradeFindBankNameAndBankCodeRequest(param)
	t.Log(data, err)
}
