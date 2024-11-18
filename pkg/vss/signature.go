package vss

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/blues120/ias-core/conf"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type VssCallbackClient struct {
	conf *conf.VssSign

	log *log.Helper
}

// NewVssCallbackClient 实例化NewVssCallbackClient
func NewVssCallbackClient(conf *conf.VssSign, logger log.Logger) *VssCallbackClient {
	return &VssCallbackClient{
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

// 生成加密数生的业务参数
func (client *VssCallbackClient) genVssParam(action, encodeString string) (string, error) {
	param, err := client.encodeVssParam(encodeString)
	if err != nil {
		return "", err
	}
	return client.createVssParam(action, param)
}

// 生成数生接口签名
func (client *VssCallbackClient) createVssParam(action, param string) (string, error) {
	nonce, err := client.generateRandomString(16)
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf("accessKey%vaction%vid%vnonce%vparams%vtimestamp%vversion%v",
		client.conf.AccessKey, action, client.conf.Id, nonce, param, timestamp, client.conf.Version)
	client.log.Debug("签名参数=" + message)
	// HMACSHA256加密
	sign := client.encryptHMACSHA256(message, client.conf.AccessSecret)
	result := fmt.Sprintf("accessKey=%v&action=%v&id=%v&nonce=%v&params=%v&timestamp=%v&version=%v&signature=%v",
		client.conf.AccessKey, action, client.conf.Id, nonce, param, timestamp, client.conf.Version, sign)
	return result, nil
}

// 加密数生的业务参数
func (client *VssCallbackClient) encodeVssParam(encodeString string) (string, error) {
	accessSecret := client.conf.AccessSecret
	key := []byte(accessSecret)
	iv := []byte(accessSecret[16:32])
	encrypted, err := client.encryptAES([]byte(encodeString), key, iv)
	if err != nil {
		log.Errorf("Encryption error: %v", err)
		return "", err
	}
	encryptedHex := hex.EncodeToString(encrypted)
	return strings.ToUpper(encryptedHex), nil
}

// EncryptAES aes加密,CBC模式，PKCS7Padding填充
func (client *VssCallbackClient) encryptAES(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = client.pkcs7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

func (client *VssCallbackClient) pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// encryptHMACSHA256 hmac-sha256加密
func (client *VssCallbackClient) encryptHMACSHA256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	encrypted := h.Sum(nil)
	return hex.EncodeToString(encrypted)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (client *VssCallbackClient) generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		// 生成一个随机的索引值
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		bytes[i] = letterBytes[index.Int64()]
	}
	return string(bytes), nil
}

// Post send request
func (client *VssCallbackClient) Post(url string, action string, payload []byte) ([]byte, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("%s%s", client.conf.Host, url)
	}
	client.log.Debugf("VssCallbackClient Post to: %+v", url)
	vssParam, err := client.genVssParam(action, string(payload))
	if err != nil {
		return nil, fmt.Errorf("VssCallbackClient: GenVssParam error: %s", err.Error())
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(vssParam)))
	if err != nil {
		return nil, fmt.Errorf("VssCallbackClient: create post request error: %s", err.Error())
	}
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	clientReq := &http.Client{}
	resp, err := clientReq.Do(req)
	if err != nil {
		return nil, fmt.Errorf("VssCallbackClient post err %+v", err.Error())
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
