package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

func CalcSHA256(input string) string {
	// 创建一个新的 SHA-256 哈希对象
	hasher := sha256.New()
	// 将输入字符串转换为字节并写入哈希对象
	hasher.Write([]byte(input))
	// 计算哈希值并返回结果
	hash := hasher.Sum(nil)
	// 将哈希值转换为十六进制格式的字符串
	return hex.EncodeToString(hash)
}
