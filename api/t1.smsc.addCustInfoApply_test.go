package api

import "testing"

func TestT1SmscAddCustInfoApplyRequest(t *testing.T) {
	param := &T1SmscAddCustInfoApplyParam{}
	data, err := sdk.T1SmscAddCustInfoApplyRequest(param)
	t.Log(data, err)
}
