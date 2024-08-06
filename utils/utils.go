package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetCurrentTimeStamp 返回当前时间的字符串，格式为 "yyyy-MM-dd HH:mm:ss"
func GetCurrentTimeStamp() string {
	now := time.Now()
	const layout = "2006-01-02 15:04:05"
	return now.Format(layout)
}

// DownloadFile 下载远程文件并返回其二进制内容
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: %s", resp.Status)
	}

	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}
