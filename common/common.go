package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"ys_sdk/utils"
)

// RequestPayload 定义公共请求参数的结构
type RequestPayload struct {
	TimeStamp  string `json:"timeStamp"`
	Method     string `json:"method"`
	Charset    string `json:"charset"`
	Sign       string `json:"sign"`
	Check      string `json:"check"`
	BizContent string `json:"bizContent"`
	ReqID      string `json:"reqId"`
	CertID     string `json:"certId"`
	Version    string `json:"version"`
}

// ResponsePayload 定义响应负载的结构
type ResponsePayload struct {
	Code         string `json:"code"`
	Msg          string `json:"msg"`
	SubCode      string `json:"subCode"`
	SubMsg       string `json:"subMsg"`
	TimeStamp    string `json:"timeStamp"`
	Norce        string `json:"norce"`
	Sign         string `json:"sign"`
	BusinessData string `json:"businessData"`
}

type Config struct {
	ReqID  string `json:"req_id"`
	CertID string `json:"cert_id"`
}

var AppConfig Config

func LoadConfig(configFile string) {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("无法打开配置文件: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(io.Reader(file))
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	if err := json.Unmarshal(bytes, &AppConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}

// 生成公共配置参数
func GenerateRequestPayload(method, bizContent string) (*ResponsePayload, error) {

	timeStamp := utils.GetCurrentTimeStamp()

	sign := "sign"

	check := "check"

	payload := RequestPayload{
		TimeStamp:  timeStamp,
		Method:     method,
		Charset:    "utf-8", // 固定值
		Sign:       sign,
		Check:      check,
		BizContent: bizContent,
		ReqID:      AppConfig.ReqID,
		CertID:     AppConfig.CertID,
		Version:    "1.0", // 固定值
	}

	url := "abc"

	return sendRequest(url, payload)
}

// sendRequest 发送HTTP请求到API
func sendRequest(url string, payload RequestPayload) (*ResponsePayload, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("JSON序列化错误: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求错误: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// 设置其他必要的头信息

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("收到非200响应: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, fmt.Errorf("读取响应体错误: %v", err)
	}

	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return nil, fmt.Errorf("JSON反序列化错误: %v", err)
	}

	return &responsePayload, nil
}
