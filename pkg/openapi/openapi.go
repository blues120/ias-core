package openapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type OpenClient struct {
	Host      string
	AppId     string
	AppSecret string
	AC        string
	BoxId     string

	log *log.Helper
}

// NewOpenClient new an open-api client
func NewOpenClient(host, appId, appSecret, ac string, boxId string, log *log.Helper) *OpenClient {
	return &OpenClient{
		Host:      host,
		AppId:     appId,
		AppSecret: appSecret,
		AC:        ac,
		BoxId:     boxId,
		log:       log,
	}
}

// Post send request
func (apiClient *OpenClient) Post(uri string, reqBody io.Reader, opts ...Option) (string, error) {
	opt := newOptions()
	for _, o := range opts {
		o.apply(&opt)
	}

	url := fmt.Sprintf("%s%s", apiClient.Host, uri)
	apiClient.log.Debugf("OpenClient Post to: %+v", url)
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return "", fmt.Errorf("OpenClient: create post request error: %s", err.Error())
	}

	timestamp, signature, err := apiClient.getSignature(req, apiClient.AppId, apiClient.AppSecret, opt.trimURIPrefix)
	if err != nil {
		apiClient.log.Debugf("OpenClient Post: getSignature params: req: %+v, appID: %s, appSecret: %s, trimURIPrefix: %+v", req, apiClient.AppId, apiClient.AppSecret, opt.trimURIPrefix)
		return "", fmt.Errorf("OpenClient: get signature error: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(opt.header.now, timestamp)
	req.Header.Add(opt.header.app, apiClient.AppId)
	req.Header.Add(opt.header.signature, signature)
	req.Header.Add(opt.header.box, apiClient.BoxId)
	if opt.header.ac != "" {
		req.Header.Add(opt.header.ac, apiClient.AC)
	}
	client := &http.Client{}
	if opt.client.transport != nil {
		client.Transport = opt.client.transport
	}

	apiClient.log.Debugf("OpenClient Post: %+v", req)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenClient post err %+v", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// Get get request
func (apiClient *OpenClient) Get(uri string, opts ...Option) (string, error) {
	opt := newOptions()
	for _, o := range opts {
		o.apply(&opt)
	}

	url := fmt.Sprintf("%s%s", apiClient.Host, uri)
	apiClient.log.Debugf("OpenClient Get to: %+v", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("OpenClient: create get request error: %s", err.Error())
	}

	timestamp, signature, err := apiClient.getSignature(req, apiClient.AppId, apiClient.AppSecret, opt.trimURIPrefix)
	if err != nil {
		return "", fmt.Errorf("OpenClient: get signature error: %s", err.Error())
	}

	req.Header.Add(opt.header.now, timestamp)
	req.Header.Add(opt.header.app, apiClient.AppId)
	req.Header.Add(opt.header.signature, signature)
	req.Header.Add(opt.header.box, apiClient.BoxId)
	if opt.header.ac != "" {
		req.Header.Add(opt.header.ac, apiClient.AC)
	}
	client := &http.Client{}
	if opt.client.transport != nil {
		client.Transport = opt.client.transport
	}

	apiClient.log.Debugf("OpenClient Get: %+v", req)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenClient get err %+v", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// SSEGet get request
func (apiClient *OpenClient) SSEGet(uri string, opts ...Option) (io.ReadCloser, error) {
	opt := newOptions()
	for _, o := range opts {
		o.apply(&opt)
	}

	url := fmt.Sprintf("%s%s", apiClient.Host, uri)
	apiClient.log.Debugf("OpenClient SSEGet to: %+v", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("OpenClient: create get request error: %s", err.Error())
	}

	// 设置SSE请求头
	req.Header.Set("Accept", "text/event-stream")

	timestamp, signature, err := apiClient.getSignature(req, apiClient.AppId, apiClient.AppSecret, opt.trimURIPrefix)
	if err != nil {
		return nil, fmt.Errorf("OpenClient: get signature error: %s", err.Error())
	}

	req.Header.Add(opt.header.now, timestamp)
	req.Header.Add(opt.header.app, apiClient.AppId)
	req.Header.Add(opt.header.signature, signature)
	req.Header.Add(opt.header.box, apiClient.BoxId)
	if opt.header.ac != "" {
		req.Header.Add(opt.header.ac, apiClient.AC)
	}
	client := &http.Client{}
	apiClient.log.Debugf("OpenClient Get: %+v", req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("OpenClient get err %+v", err.Error())
	}
	// defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	eventStream := resp.Body
	return eventStream, nil
}

func hmacSha256Byte(target, key string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(target))
	hashBytes := h.Sum(nil)
	return hashBytes
}

func encrypt(content, key string) (signature string, err error) {
	// 替换空格为+
	key = strings.ReplaceAll(key, " ", "+")
	// 替换-为+号
	key = strings.ReplaceAll(key, "-", "+")
	// 替换_为/号
	key = strings.ReplaceAll(key, "_", "/")
	// 填充=，字节为4的倍数
	for len(key)%4 != 0 {
		key += "="
	}
	bcode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return
	}

	sigbyte := hmacSha256Byte(content, string(bcode))
	sigstr := base64.URLEncoding.EncodeToString(sigbyte)
	signature = strings.Replace(sigstr, "=", "", -1)
	return
}

func (apiClient *OpenClient) getSignature(req *http.Request, appID string, appSecret string, opts ...Option) (string, string, error) {
	opt := newOptions()
	for _, o := range opts {
		o.apply(&opt)
	}
	apiClient.log.Debugf("getSignature: opts: %+v", opt)

	uri := req.URL.RequestURI()
	if opt.trimURIPrefix.prefix != "" { // 去除前缀
		uri = strings.TrimPrefix(uri, opt.trimURIPrefix.prefix)
	}
	apiClient.log.Debugf("getSignature: uri before %s after: %s", req.URL.RequestURI(), uri)

	timestamp_ms := time.Now().Unix() * 1000
	timestamp_day := timestamp_ms / 86400000
	timestamp_ms_str := strconv.FormatInt(timestamp_ms, 10)
	sign_str := fmt.Sprintf("%s\n%v\n%s", appID, timestamp_ms, uri)
	identity := fmt.Sprintf("%s:%v", appID, timestamp_day)
	apiClient.log.Debugf("getSignature: sign_str: %s, identity: %s", sign_str, identity)

	tmp_signature, err := encrypt(identity, appSecret)
	if err != nil {
		return "", "", fmt.Errorf("getSignature: encrypt identity failed, err:%v", err)
	}

	apiClient.log.Debugf("getSignature: tmp_signature: %s", tmp_signature)
	signature, err := encrypt(sign_str, tmp_signature)
	if err != nil {
		return "", "", fmt.Errorf("getSignature: encrypt sign_str failed, err:%v", err)
	}

	return timestamp_ms_str, signature, nil
}
