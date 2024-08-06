package api

import (
	"log"
	"ys_sdk/common"
	"ys_sdk/utils"
)

// GenerateConfig 生成4.1配置
func GenerateFileSmscUpload(bizContent string) common.RequestPayload {
	method := "file.smsc.upload"
	return common.GenerateRequestPayload(method, bizContent)
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

func GenerateFileSmscUploadContent() FileSmcsUploadContent {
	fileURL := "http://example.com/path/to/your/file.jpg"
	fileData, err := utils.DownloadFile(fileURL)
	if err != nil {
		log.Fatalf("无法下载文件: %v", err)
	}

	meta := GenerateFileSmscUploadMeta("a", "b", "c", "d")

	return FileSmcsUploadContent{
		File: fileData,
		Meta: meta,
	}

}

func GenerateFileSmscUploadMeta(sha256, picType, picName, sysFlowID string) FileSmcsUploadMeta {
	return FileSmcsUploadMeta{
		Sha256:    sha256,
		PicType:   picType,
		PicNm:     picName,
		SysFlowID: sysFlowID,
	}
}
