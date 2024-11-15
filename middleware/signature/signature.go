package signature

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
)

var (
	ErrSignature = errors.New("请求签名错误")
)

// NewSignatureMiddleware 验签中间件
func NewSignatureMiddleware(sigUc *biz.SignatureUsecase, opts ...Option) middleware.Middleware {
	options := NewOptions()
	for _, o := range opts {
		o.Apply(&options)
	}
	helper := log.NewHelper(log.With(log.DefaultLogger, "module", "[SignatureMiddleware]"))

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			r, ok := http.RequestFromServerContext(ctx)
			if ok {
				// 从请求头中获取签名信息
				timestamp := r.Header.Get(options.H.now)
				appId := r.Header.Get(options.H.app)
				signature := r.Header.Get(options.H.signature)
				boxId := r.Header.Get(options.H.box)
				if timestamp == "" || appId == "" || signature == "" {
					helper.Debugf("verifySignature params: timestamp:%s,appId:%s,signature:%s", timestamp, appId, signature)
					return nil, fmt.Errorf("请求签名错误: 字段为空")
				}

				// 获取对应的 appSecret
				sig, err := sigUc.FindByCondition(ctx, boxId, appId)
				if err != nil {
					helper.Debugf("verifySignature params: boxId:%s,appId:%s,signature:%s", boxId, appId, signature)
					return nil, fmt.Errorf("请求签名错误: 找不到boxId或appId")
				}

				// 验签
				if err := verifySignature(r.RequestURI, timestamp, appId, sig.AppSecret, signature); err != nil {
					helper.Debugf("verifySignature params: uri:%s, timestamp:%s, appId:%s, appSecret:%s, signature:%s", r.RequestURI, timestamp, appId, sig.AppSecret, signature)
					return nil, fmt.Errorf("请求签名错误: 验签出错 %v", err)
				}
				return handler(ctx, req)
			}
			return nil, ErrSignature
		}
	}
}

func verifySignature(uri string, timestamp, appID, appSecret, signature string) error {
	timestampMs, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("verifySignature: timestampMs parse failed, err:%v", err)
	}
	timestampDay := timestampMs / 86400000
	identity := fmt.Sprintf("%s:%v", appID, timestampDay)
	tmpSignature, err := encrypt(identity, appSecret)
	if err != nil {
		return fmt.Errorf("verifySignature: encrypt identity failed, err:%v", err)
	}

	signStr := fmt.Sprintf("%s\n%v\n%s", appID, timestampMs, uri)
	formalSignature, err := encrypt(signStr, tmpSignature)
	if err != nil {
		return fmt.Errorf("verifySignature: encrypt sign_str failed, err:%v", err)
	}
	if formalSignature != signature {
		return fmt.Errorf("verifySignature: signature not match: formalSignature[%s] signature[%s]", formalSignature, signature)
	}

	return nil
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
	b64Str, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return
	}
	sigByte := hmacSha256Byte(content, string(b64Str))
	sigStr := base64.URLEncoding.EncodeToString(sigByte)
	signature = strings.Replace(sigStr, "=", "", -1)
	return
}
