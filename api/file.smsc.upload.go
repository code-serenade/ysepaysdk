package api

import (
	"log"

	"github.com/code-serenade/ysepaysdk/utils"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// GenerateConfig 生成4.1配置
func (c *Config) RequstFileSmscUpload(filePath, picType, sysFlowID string) (data xmap.M, err error) {
	method := "file.smsc.upload"
	version := "1.0"
	url := methodToUrl(method)
	bizContent := generateFileSmscUploadContent(filePath, picType, sysFlowID)
	_, data, err = c.UploadRequest(url, method, version, filePath, bizContent)
	return
}

// Meta 定义媒体文件元信息
type FileSmcsUploadMeta struct {
	Sha256    string `json:"sha256,omitempty"`
	PicType   string `json:"picType,omitempty"`
	PicNm     string `json:"picNm,omitempty"`
	SysFlowID string `json:"sysFlowId,omitempty"`
}

// BizContent 定义业务请求参数
type FileSmcsUploadContent struct {
	// File []byte             `json:"file,omitempty"`
	Meta FileSmcsUploadMeta `json:"meta,omitempty"`
}

func generateFileSmscUploadContent(filePath, picType, sysFlowID string) string {
	fileData, err := utils.ReadLocalFile(filePath)
	if err != nil {
		log.Fatalf("无法下载文件: %v", err)
	}
	hash := utils.CalcSHA256(string(fileData))
	meta := generateFileSmscUploadMeta(hash, picType, "", sysFlowID)

	content := FileSmcsUploadContent{
		// File: fileData,
		Meta: meta,
	}
	return converter.JSON(content)
}

func generateFileSmscUploadMeta(sha256, picType, picName, sysFlowID string) FileSmcsUploadMeta {
	return FileSmcsUploadMeta{
		Sha256:    sha256,
		PicType:   picType,
		PicNm:     picName,
		SysFlowID: sysFlowID,
	}
}
