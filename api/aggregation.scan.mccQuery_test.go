package api

import (
	"testing"

	"github.com/codingeasygo/util/converter"
)

func TestAggregationScanMccQueryRequest(t *testing.T) {
	// param := NewAggregationScanMccQueryParam("2")
	param := &AggregationScanMccQueryParam{
		MccCd: "7372",
	}
	data, _ := sdk.AggregationScanMccQueryRequest(param)
	t.Log(converter.JSON(data))
}
