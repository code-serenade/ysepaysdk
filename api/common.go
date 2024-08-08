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
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/code-serenade/easycrypto"
	"github.com/code-serenade/ysepaysdk/utils"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

// RequestPayload 定义公共请求参数的结构
type RequestPayload struct {
	TimeStamp  string `json:"timeStamp,omitempty"`
	Method     string `json:"method,omitempty"`
	Charset    string `json:"charset,omitempty"`
	Sign       string `json:"sign,omitempty"`
	Check      string `json:"check,omitempty"`
	BizContent string `json:"bizContent,omitempty"`
	ReqID      string `json:"reqId,omitempty"`
	CertID     string `json:"certId,omitempty"`
	Version    string `json:"version,omitempty"`
}

// NewRequestPayload 创建新的请求负载
func NewRequestPayload(method, version string) *RequestPayload {
	return &RequestPayload{
		ReqID:     utils.CurrentYYMMDDHHMMSSS(),
		TimeStamp: utils.GetCurrentTimeStamp(),
		Method:    method,
		Charset:   "utf-8", // 固定值
		Version:   version,
	}
}

// EncryptCheck 使用公钥和AES密钥加密Check字段
func (r *RequestPayload) EncryptCheck(pubKey, aesKey []byte) error {
	check, err := easycrypto.RSAEncrypt(pubKey, aesKey)
	if err != nil {
		return err
	}
	r.Check = base64.StdEncoding.EncodeToString(check)
	return nil
}

// EncryptBizContent 使用AES密钥加密业务内容
func (r *RequestPayload) EncryptBizContent(keyByte []byte) error {
	encryptedContent, err := easycrypto.AESEncryptECB(r.BizContent, keyByte)
	if err != nil {
		return err
	}
	r.BizContent = encryptedContent
	return nil
}

// makeSignBefore 生成签名前的字符串
func (r *RequestPayload) makeSignBefore() string {
	m := xmap.M{
		"timeStamp":  r.TimeStamp,
		"method":     r.Method,
		"charset":    r.Charset,
		"reqId":      r.ReqID,
		"certId":     r.CertID,
		"version":    r.Version,
		"check":      r.Check,
		"bizContent": r.BizContent,
	}
	return utils.MapToUrlValues(m)
}

// CalcSign 计算签名
func (r *RequestPayload) CalcSign(key []byte) error {
	content := r.makeSignBefore()

	if Verbose {
		log.Printf("before CalcSign string %v", content)
	}

	sign, err := easycrypto.RSASign(key, []byte(content))
	if err != nil {
		return err
	}
	r.Sign = sign
	return nil
}

// EncodeMap 编码请求负载为map
func (r *RequestPayload) EncodeMap() map[string]string {
	return map[string]string{
		"timeStamp":  r.TimeStamp,
		"method":     r.Method,
		"charset":    r.Charset,
		"sign":       r.Sign,
		"check":      r.Check,
		"bizContent": r.BizContent,
		"reqId":      r.ReqID,
		"certId":     r.CertID,
		"version":    r.Version,
	}
}

// ResponsePayload 定义响应负载的结构
type ResponsePayload struct {
	Code         string `json:"code,omitempty"`
	Msg          string `json:"msg,omitempty"`
	SubCode      string `json:"subCode,omitempty"`
	SubMsg       string `json:"subMsg,omitempty"`
	TimeStamp    string `json:"timeStamp,omitempty"`
	Norce        string `json:"norce,omitempty"`
	Sign         string `json:"sign,omitempty"`
	BusinessData string `json:"businessData,omitempty"`
}

func (r *ResponsePayload) EncodeMap() (m xmap.M) {
	m = xmap.M{
		"code":         r.Code,
		"msg":          r.Msg,
		"subCode":      r.SubCode,
		"subMsg":       r.SubMsg,
		"timeStamp":    r.TimeStamp,
		"norce":        r.Norce,
		"sign":         r.Sign,
		"businessData": r.BusinessData,
	}
	return
}

// Config 配置结构体
type Config struct {
	CertID     string `json:"cert_id,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
	PublicKey  string `json:"public_key,omitempty"`
}

// NewConfig 创建新的配置
func NewConfig(conf xmap.M) *Config {
	return &Config{
		CertID:     conf.Str("cert_id"),
		PrivateKey: conf.Str("private_key"),
		PublicKey:  conf.Str("public_key"),
	}
}

func NewConfigFromFile(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

// Decode 解码业务数据
func (c *Config) Decode(aesKey []byte, businessData string) (xmap.M, error) {
	if len(businessData) < 1 {
		return nil, fmt.Errorf("businessData is empty")
	}
	decryptedBizData, err := easycrypto.AESDecryptECB(businessData, aesKey)
	if err != nil {
		return nil, err
	}
	var data xmap.M
	_, err = converter.UnmarshalJSON(bytes.NewBuffer(decryptedBizData), &data)
	return data, err
}

// Request 发送普通POST请求
func (c *Config) Request(url, method, version, bizContent string) (*ResponsePayload, xmap.M, error) {
	if Verbose {
		log.Printf("request bizContent %v", bizContent)
	}
	payload := NewRequestPayload(method, version)
	payload.CertID = c.CertID
	aesKey := []byte(getRandomString(16))
	// 加密check
	err := payload.EncryptCheck([]byte(c.PublicKey), aesKey)
	if err != nil {
		log.Printf("加密check失败: %v", err)
		return nil, nil, err
	}
	// 加密bizContent
	payload.BizContent = bizContent
	err = payload.EncryptBizContent(aesKey)
	if err != nil {
		log.Printf("加密bizContent失败: %v", err)
		return nil, nil, err
	}
	// 处理签名
	err = payload.CalcSign([]byte(c.PrivateKey))
	if err != nil {
		log.Printf("签名失败: %v", err)
		return nil, nil, err
	}

	response, err := sendRequest(url, payload)
	if err != nil {
		return nil, nil, err
	}
	if Verbose {
		log.Printf("response %v", response)
	}
	if response.Code != successCode {
		return response, nil, fmt.Errorf("code:%s msg:%s", response.Code, response.Msg)
	}
	var result = response.EncodeMap()
	data, kerr := c.Decode(aesKey, response.BusinessData)
	if kerr == nil {
		result["bizContent"] = data
	}
	return response, result, err
}

// UploadRequest 发送文件上传请求
func (c *Config) UploadRequest(url, method, version, filePath, bizContent string) (*ResponsePayload, xmap.M, error) {
	if Verbose {
		log.Printf("request bizContent %v", bizContent)
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("打开文件失败: %v", err)
		return nil, nil, err
	}
	defer file.Close()
	payload := NewRequestPayload(method, version)
	payload.CertID = c.CertID
	aesKey := []byte(getRandomString(16))
	// 加密check
	err = payload.EncryptCheck([]byte(c.PublicKey), aesKey)
	if err != nil {
		log.Printf("加密check失败: %v", err)
		return nil, nil, err
	}
	// 加密bizContent
	payload.BizContent = bizContent
	err = payload.EncryptBizContent(aesKey)
	if err != nil {
		log.Printf("加密bizContent失败: %v", err)
		return nil, nil, err
	}
	// 处理签名
	err = payload.CalcSign([]byte(c.PrivateKey))
	if err != nil {
		log.Printf("签名失败: %v", err)
		return nil, nil, err
	}

	response, err := sendUploadRequest(url, payload, file)
	if err != nil {
		return nil, nil, err
	}
	if Verbose {
		log.Printf("response %v", converter.JSON(response))
	}
	if response.Code != successCode {
		return response, response.EncodeMap(), fmt.Errorf("code:%s msg:%s", response.Code, response.Msg)
	}
	var result, data xmap.M
	result = response.EncodeMap()
	if response.SubCode == successCode {
		_, kerr := converter.UnmarshalJSON(bytes.NewBuffer([]byte(response.BusinessData)), &data)
		if kerr == nil {
			result["bizContent"] = data
		}
	}
	return response, result, err
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("收到非200响应: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体错误: %v", err)
	}

	if Verbose {
		log.Printf("response body %v", string(body))
	}

	kbody, kerr := base64.StdEncoding.DecodeString(string(body))
	if kerr == nil {
		body = kbody
	}
	if Verbose {
		log.Printf("response body %v", string(body))
	}

	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return nil, fmt.Errorf("JSON反序列化错误: %v", err)
	}

	return &responsePayload, nil
}

// getRandomString 生成指定长度的随机字符串
func getRandomString(length int) string {
	const ALLCHAR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(ALLCHAR))))
		sb.WriteByte(ALLCHAR[num.Int64()])
	}
	return sb.String()
}

// sendUploadRequest 发送文件上传请求到API
func sendUploadRequest(url string, payload *RequestPayload, file *os.File) (*ResponsePayload, error) {
	params := payload.EncodeMap()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	fileWriter, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, fmt.Errorf("创建文件字段错误: %v", err)
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, fmt.Errorf("写入文件字段错误: %v", err)
	}

	for key, val := range params {
		err = writer.WriteField(key, val)
		if err != nil {
			return nil, fmt.Errorf("设置字段 %s 错误: %v", key, err)
		}
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("关闭 writer 错误: %v", err)
	}

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, fmt.Errorf("创建请求错误: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("收到非200响应: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体错误: %v", err)
	}

	body, _ = base64.StdEncoding.DecodeString(string(body))

	if Verbose {
		log.Printf("response body %v", string(body))
	}

	var responsePayload ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return nil, fmt.Errorf("JSON反序列化错误: %v", err)
	}

	return &responsePayload, nil
}

func mergeMap(m1, m2 xmap.M) xmap.M {
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}
