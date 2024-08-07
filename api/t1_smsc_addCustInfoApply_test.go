package api

import "testing"

func TestT1SmscAddCustInfoApplyRequest(t *testing.T) {
	data, err := sdk.T1SmscAddCustInfoApplyRequest()
	t.Log(data, err)
}
