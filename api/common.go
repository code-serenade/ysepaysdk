package api

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
	"ys_sdk/utils"

	"github.com/CodeSerenade/easycrypto"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
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

func NewRequestPayload(method, version string) *RequestPayload {
	return &RequestPayload{
		ReqID:     utils.CurrentYYMMDDHHMMSSS(),
		TimeStamp: utils.GetCurrentTimeStamp(),
		Method:    method,
		Charset:   "utf-8", // 固定值
		Version:   version,
	}
}

func (r *RequestPayload) EncryptCheck(pubKey, aesKey []byte) (err error) {
	check, err := easycrypto.RSAEncrypt(pubKey, aesKey)
	if err != nil {
		return
	}
	r.Check = base64.StdEncoding.EncodeToString(check)
	return
}

func (r *RequestPayload) EncryptBizContent(keyByte []byte) (err error) {
	r.BizContent, err = easycrypto.AESEncryptECB(r.BizContent, keyByte)
	return
}

func (r *RequestPayload) CalcSign(key []byte) (err error) {
	params := url.Values{}
	params.Set("timeStamp", r.TimeStamp)
	params.Set("method", r.Method)
	params.Set("charset", r.Charset)
	params.Set("check", r.Check)

	content := strings.ReplaceAll(params.Encode(), "+", "%20")

	r.Sign, err = easycrypto.RSASign(key, []byte(content))
	if err != nil {
		return
	}
	return
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
	// Resp         any // TODO
}

type Config struct {
	// ReqID      string `json:"req_id"`
	CertID     string `json:"cert_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func NewConfig(conf xmap.M) *Config {
	return &Config{
		// ReqID:      conf.Str("req_id"),
		CertID:     conf.Str("cert_id"),
		PrivateKey: conf.Str("private_key"),
		PublicKey:  conf.Str("public_key"),
	}
}

func (c *Config) Decode(aseKey []byte, businessData string) (data xmap.M, err error) {
	decryptedBizData, err := easycrypto.AESDecryptECB(businessData, aseKey)
	if err != nil {
		return
	}
	_, err = converter.UnmarshalJSON(bytes.NewBuffer(decryptedBizData), &data)
	return
}

func (c *Config) Request(url, method, bizContent string) (resp *ResponsePayload, data xmap.M, err error) {
	payload := NewRequestPayload(method, "1.0")
	payload.CertID = c.CertID
	aesKey := []byte(getRandomString(16))
	// 加密check
	payload.EncryptCheck([]byte(c.PublicKey), aesKey)
	// 加密bizContent
	payload.EncryptBizContent(aesKey)
	// 处理签名
	err = payload.CalcSign([]byte(c.PrivateKey))
	if err != nil {
		log.Printf("签名失败: %v", err)
		return
	}

	response, err := sendRequest(url, payload)
	if err != nil {
		return
	}
	if Verbose {
		log.Printf("response %v", response)
	}
	if response.Code != successCode {
		err = fmt.Errorf("code:%s msg:%s", response.Code, response.Msg)
		return
	}
	data, err = c.Decode(aesKey, response.BusinessData)
	return
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

// sendRequest 发送HTTP请求到API
func sendRequest(url string, payload *RequestPayload) (*ResponsePayload, error) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("JSON序列化错误: %v", err)
	}

	if Verbose {
		log.Printf("request url %v", url)
		log.Printf("request payload %v", string(jsonData))
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

func getRandomString(length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(ALLCHAR))))
		sb.WriteByte(ALLCHAR[num.Int64()])
	}
	return sb.String()
}
