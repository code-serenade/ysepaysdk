package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/big"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/codingeasygo/util/xmap"
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

// make a map to url.Values
func MapToUrlValues(m xmap.M) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	for _, k := range keys {
		if v := m.Str(k); len(v) > 0 {
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(v)
			sb.WriteString("&")
		}
	}
	signDataStr := sb.String()
	if len(signDataStr) > 0 {
		signDataStr = signDataStr[:len(signDataStr)-1]
	}
	return signDataStr
}

func MapToUrlValuesString(m xmap.M) string {
	values := url.Values{}
	for k := range m {
		if v := m.Str(k); len(v) > 0 {
			values.Set(k, v)
		}
	}
	return strings.ReplaceAll(values.Encode(), "+", "%20")
}

// GetRandomString 生成指定长度的随机字符串
func GetRandomString(length int) string {
	const ALLCHAR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(ALLCHAR))))
		sb.WriteByte(ALLCHAR[num.Int64()])
	}
	return sb.String()
}
