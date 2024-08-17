package api

import (
	"fmt"
	"testing"
	"time"
)

func TestRequstFileSmscUpload_ValidInputs(t *testing.T) {
	filePath := "../testdata/test.jpeg"
	picType := "A001"
	sysFlowID := fmt.Sprintf("%v", time.Now().UnixNano())
	data, err := sdk.RequstFileSmscUpload(filePath, picType, sysFlowID)
	t.Log(data, err)
}
