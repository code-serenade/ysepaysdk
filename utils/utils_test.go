package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/codingeasygo/util/uuid"
)

func TestGetCurrentTimeStamp(t *testing.T) {
	timestamp := GetCurrentTimeStamp()

	// 解析返回的时间戳
	const layout = "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		t.Errorf("时间戳格式错误: %v", err)
	}

	// 检查时间戳是否为当前时间
	now := time.Now()
	if parsedTime.After(now) || parsedTime.Before(now.Add(-1*time.Minute)) {
		t.Logf("时间戳不是当前时间: %v", timestamp)
	}
}

func TestUUID(t *testing.T) {
	fmt.Println(uuid.New())
}
