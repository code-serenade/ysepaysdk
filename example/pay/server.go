package main

import (
	"fmt"
	"log"
	"time"

	ysepaysdk "github.com/code-serenade/ysepaysdk/api"
)

func main() {
	sdk, err := ysepaysdk.NewConfigFromFile("../../testdata/conf.json")
	if err != nil {
		panic(err)
	}
	log.Println(sdk)
	filePath := "../../testdata/test.jpeg"
	picType := "A001"
	sysFlowID := fmt.Sprintf("%v", time.Now().UnixNano())
	data, err := sdk.RequstFileSmscUpload(filePath, picType, sysFlowID)
	log.Println(data)
	if err != nil {
		log.Println(err)
		return
	}
}
