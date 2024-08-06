package api

import (
	"log"
	"ys_sdk/common"
	"ys_sdk/utils"
)

// GenerateConfig 生成4.1配置
func RequstFileSmscUpload(filePath, picType, sysFlowID string) {
	method := "file.smsc.upload"

	bizContent := generateFileSmscUploadContent(filePath, picType, sysFlowID)
	respon, err := common.GenerateRequestPayload(method, bizContent)

	if err != nil {
		// TODO respon
		print(respon)
	}
}

// Meta 定义媒体文件元信息
type FileSmcsUploadMeta struct {
	Sha256    string `json:"sha256"`
	PicType   string `json:"picType"`
	PicNm     string `json:"picNm"`
	SysFlowID string `json:"sysFlowId"`
}

// BizContent 定义业务请求参数
type FileSmcsUploadContent struct {
	File []byte             `json:"file"`
	Meta FileSmcsUploadMeta `json:"meta"`
}

func generateFileSmscUploadContent(filePath, picType, sysFlowID string) string {
	fileData, err := utils.ReadLocalFile(filePath)
	if err != nil {
		log.Fatalf("无法下载文件: %v", err)
	}

	meta := generateFileSmscUploadMeta("a", "b", "c", "d")

	content := FileSmcsUploadContent{
		File: fileData,
		Meta: meta,
	}

}

func generateFileSmscUploadMeta(sha256, picType, picName, sysFlowID string) FileSmcsUploadMeta {
	return FileSmcsUploadMeta{
		Sha256:    sha256,
		PicType:   picType,
		PicNm:     picName,
		SysFlowID: sysFlowID,
	}
}
