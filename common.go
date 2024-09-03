package ysepaysdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/code-serenade/easycrypto"
	"github.com/code-serenade/ysepaysdk/utils"
	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xmap"
)

type APIRequest struct {
	Method     string `json:"method"`               // 接口名称，例如ysepay.online.qrcodepay
	PartnerID  string `json:"partner_id"`           // 服务商商户号
	Timestamp  string `json:"timestamp"`            // 请求时间，格式"yyyy-MM-dd HH:mm:ss"
	Charset    string `json:"charset"`              // 编码格式，例如UTF-8, GBK, GB2312
	SignType   string `json:"sign_type"`            // 签名算法，例如RSA、SM
	Sign       string `json:"sign"`                 // 签名字符串，Base64编码
	NotifyURL  string `json:"notify_url"`           // 异步通知地址
	Version    string `json:"version"`              // API版本号
	ReturnURL  string `json:"return_url,omitempty"` // 同步通知地址，可选
	TranType   string `json:"tran_type,omitempty"`  // 交易类型，可选
	BizContent string `json:"biz_content"`          // 业务请求参数集合
}

func NewAPIRequest(method, partnerID, notifyURL, bizContent string) *APIRequest {
	return &APIRequest{
		Method:     method,
		PartnerID:  partnerID,
		Timestamp:  time.Now().Format(`2006-01-02 15:04:05`),
		Charset:    "UTF-8",
		SignType:   "RSA",
		Version:    "3.5",
		NotifyURL:  notifyURL,
		BizContent: bizContent,
	}
}

func (r *APIRequest) Map() xmap.M {
	m := xmap.M{}
	converter.UnmarshalJSON(bytes.NewBufferString(converter.JSON(r)), &m)
	return m
}

type APIResponse xmap.M

// var Shared *Config

type Config map[string]ConfigItem

type ConfigItem struct {
	PartnerID     string `json:"partner_id"`      // 商户号
	YSPublicKey   string `json:"ys_public_key"`   // 银盛公钥
	MerPrivateKey string `json:"mer_private_key"` // 商户私钥
}

func NewConfigFromFile(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func NewConfigFromItem(partnerID, ysPublicKey, merPrivateKey string) Config {
	config := map[string]ConfigItem{}
	config[partnerID] = ConfigItem{
		PartnerID:     partnerID,
		YSPublicKey:   ysPublicKey,
		MerPrivateKey: merPrivateKey,
	}
	return config
}

func (c Config) MerPrivateKey(partnerID string) string {
	return c[partnerID].MerPrivateKey
}
func (c Config) YSPublicKey(partnerID string) string {
	return c[partnerID].YSPublicKey
}

func (c Config) Request(apiUrl string, request *APIRequest) (resp xmap.M, err error) {
	signContent := utils.MapToUrlValues(request.Map())
	if Verbose {
		log.Printf("Request(%v) sign content %v", request.PartnerID, signContent)
	}
	privKey := c.MerPrivateKey(request.PartnerID)
	sign, err := easycrypto.RSASignWithSHA1([]byte(privKey), []byte(signContent))
	if err != nil {
		if Verbose {
			log.Printf("Request(%v) privKey %v\n", request.PartnerID, privKey)
		}
		return nil, err
	}
	request.Sign = sign
	return sendRequest(apiUrl, request)
}

func sendRequest(url string, request *APIRequest) (response xmap.M, err error) {
	payload := utils.MapToUrlValuesString(request.Map())
	if Verbose {
		log.Printf("Request(%v) url %v", request.PartnerID, url)
		log.Printf("Request(%v) payload %v", request.PartnerID, payload)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("user-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;SV1)")
	// 发送请求
	httpResp, err := client.Do(req)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if Verbose {
		body, _ := io.ReadAll(httpResp.Body)
		log.Printf("Request(%v) response %v", request.PartnerID, string(body))
	}

	err = json.NewDecoder(httpResp.Body).Decode(&response)
	if err != nil {
		return
	}
	return response, nil
}
