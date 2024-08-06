package utils

import (
	"io"
	"os"
	"time"
)

// GetCurrentTimeStamp 返回当前时间的字符串，格式为 "yyyy-MM-dd HH:mm:ss"
func GetCurrentTimeStamp() string {
	now := time.Now()
	const layout = "2006-01-02 15:04:05"
	return now.Format(layout)
}

func CurrentYYMMDDHHMMSSS() string {
	now := time.Now()
	const layout = "20060102150405"
	return now.Format(layout)
}

func ReadLocalFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}
