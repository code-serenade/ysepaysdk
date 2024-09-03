package api

import (
	"log"

	"github.com/code-serenade/ysepaysdk/utils"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// GenerateConfig 生成4.1配置
func (c *Config) FileSmscUploadChangePicRequest(filePath, picType, sysFlowID string) (data xmap.M, err error) {
	method := "file.smsc.uploadChangePic"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := generateFileSmscUploadChangePicContent(filePath, picType, sysFlowID)
	_, data, err = c.UploadRequest(url, method, version, filePath, bizContent)
	return
}

// Meta 定义媒体文件元信息
type FileSmcsUploadChangePicMeta struct {
	Sha256       string `json:"sha256,omitempty"`
	PicType      string `json:"picType,omitempty"`
	PicNm        string `json:"picNm,omitempty"`
	ChangeFlowId string `json:"changeFlowId,omitempty"`
}

// BizContent 定义业务请求参数
type FileSmcsUploadChangePicContent struct {
	Meta FileSmcsUploadChangePicMeta `json:"meta,omitempty"`
}

func generateFileSmscUploadChangePicContent(filePath, picType, changeFlowId string) string {
	fileData, err := utils.ReadLocalFile(filePath)
	if err != nil {
		log.Fatalf("无法下载文件: %v", err)
	}
	hash := utils.CalcSHA256(string(fileData))
	meta := generateFileSmscUploadChangePicMeta(hash, picType, "", changeFlowId)

	content := FileSmcsUploadChangePicContent{
		Meta: meta,
	}
	return converter.JSON(content)
}

// File: fileData,xx
func generateFileSmscUploadChangePicMeta(sha256, picType, picName, sysFlowID string) FileSmcsUploadChangePicMeta {
	return FileSmcsUploadChangePicMeta{
		Sha256:       sha256,
		PicType:      picType,
		PicNm:        picName,
		ChangeFlowId: sysFlowID,
	}
}
