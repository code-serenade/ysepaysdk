package api

import "testing"

func TestSmscqueryCustApplyRequest(t *testing.T) {
	sdk, _ := NewConfigFromFile("../temp.json")
	IsDev = false
	param := NewSmscqueryCustApplyParam("APPL202408091053204400181", "")
	data, err := sdk.SmscqueryCustApplyRequest(param)
	t.Log(data, err)
}
