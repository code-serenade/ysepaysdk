package api

import "testing"

func TestT1SmscAuditCustInfoApplyRequest(t *testing.T) {

	param := NewT1SmscAuditCustInfoApplyParam("入网申请流水号")

	data, err := sdk.T1SmscAuditCustInfoApplyRequest(param)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

}