package api

import "testing"

func TestT1SmscAddCustInfoApplyRequest(t *testing.T) {
	IsDev = false
	data, err := sdk.T1SmscAddCustInfoApplyRequest()
	if err != nil {
		t.Errorf("TestT1SmscAddCustInfoApplyRequest failed: %v", err)
		return
	}
	t.Logf("TestT1SmscAddCustInfoApplyRequest success: %v", data)

}
