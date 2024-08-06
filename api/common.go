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
	"net/url"
	"os"
	"sort"
	"strings"
	"ys_sdk/utils"

	"github.com/CodeSerenade/easycrypto"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xhttp"
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

func (r *RequestPayload) makeSignBefore() string {
	m := xmap.M{}
	// converter.UnmarshalJSON(bytes.NewBuffer([]byte(converter.JSON(r))), &m)
	m["timeStamp"] = r.TimeStamp
	m["method"] = r.Method
	m["charset"] = r.Charset
	m["reqId"] = r.ReqID
	m["certId"] = r.CertID
	m["version"] = r.Version
	m["check"] = r.Check
	m["bizContent"] = r.BizContent

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(m.Str(k))
		sb.WriteString("&")
	}
	signDataStr := sb.String()
	if len(signDataStr) > 0 {
		signDataStr = signDataStr[:len(signDataStr)-1]
	}
	return signDataStr
}

func (r *RequestPayload) CalcSign(key []byte) (err error) {
	content := r.makeSignBefore()
	// content := strings.ReplaceAll(params.Encode(), "+", "%20")

	if Verbose {
		log.Printf("before CalcSign string %v", content)
	}

	r.Sign, err = easycrypto.RSASign(key, []byte(content))
	if err != nil {
		return
	}
	return
}

type RequestUploadPayload struct {
	RequestPayload
	File string `json:"file"`
}

func (r *RequestUploadPayload) SetFile(file string) {
	r.File = file
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

func (c *Config) Request(url, method, version, bizContent string) (resp *ResponsePayload, data xmap.M, err error) {
	if Verbose {
		log.Printf("request bizContent %v", bizContent)
	}
	payload := NewRequestPayload(method, version)
	payload.CertID = c.CertID
	aesKey := []byte(getRandomString(16))
	// 加密check
	payload.EncryptCheck([]byte(c.PublicKey), aesKey)
	// 加密bizContent
	payload.BizContent = bizContent
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

func (c *Config) UploadRequest(url, method, version, fileData, bizContent string) (resp *ResponsePayload, data xmap.M, err error) {
	if Verbose {
		log.Printf("request bizContent %v", bizContent)
	}
	payload := NewRequestPayload(method, version)
	payload.CertID = c.CertID
	aesKey := []byte(getRandomString(16))
	// 加密check
	payload.EncryptCheck([]byte(c.PublicKey), aesKey)
	// 加密bizContent
	payload.BizContent = bizContent
	payload.EncryptBizContent(aesKey)
	// 处理签名
	err = payload.CalcSign([]byte(c.PrivateKey))
	if err != nil {
		log.Printf("签名失败: %v", err)
		return
	}

	uploadPayload := &RequestUploadPayload{
		RequestPayload: *payload,
		File:           fileData,
	}

	response, err := sendUploadRequest(url, uploadPayload)
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

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("收到非200响应: %v", resp.StatusCode)
	// }

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

func getRandomString(length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(ALLCHAR))))
		sb.WriteByte(ALLCHAR[num.Int64()])
	}
	return sb.String()
}

func (r *RequestUploadPayload) Encode() string {
	params := url.Values{}
	params.Set("timeStamp", r.TimeStamp)
	params.Set("method", r.Method)
	params.Set("charset", r.Charset)
	params.Set("check", r.Check)
	params.Set("sign", r.Sign)
	params.Set("bizContent", r.BizContent)
	params.Set("reqId", r.ReqID)
	params.Set("certId", r.CertID)
	params.Set("version", r.Version)
	return params.Encode()
}

// sendRequest 发送HTTP请求到API
func sendUploadRequest(url string, payload *RequestUploadPayload) (*ResponsePayload, error) {
	// fileds.SetValue("file", payload.File)
	result, kerr := xhttp.UploadMap(nil, "file", payload.File, "%v?%v", url, payload.Encode())

	if Verbose {
		log.Printf("response body %v kerr %v", converter.JSON(result), kerr)
	}

	// 创建一个缓冲区用来存放multipart/form-data的内容
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	writer.WriteField("timeStamp", payload.TimeStamp)
	writer.WriteField("method", payload.Method)
	writer.WriteField("charset", payload.Charset)
	writer.WriteField("check", payload.Check)
	writer.WriteField("sign", payload.Sign)
	writer.WriteField("bizContent", payload.BizContent)
	writer.WriteField("reqId", payload.ReqID)
	writer.WriteField("certId", payload.CertID)
	writer.WriteField("version", payload.Version)
	writer.WriteField("file", payload.File)

	// 关闭writer以完成multipart/form-data的写入
	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("关闭 writer 错误: %v", err)
	}

	if Verbose {
		log.Printf("request url %v", url)
		log.Printf("request payload %v", converter.JSON(payload))
	}

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, fmt.Errorf("创建请求错误: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	// 设置其他必要的头信息

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("收到非200响应: %v", resp.StatusCode)
	// }

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, fmt.Errorf("读取响应体错误: %v", err)
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
