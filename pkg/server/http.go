package server

import (
	"encoding/json"
	nethttp "net/http"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	baseContentType = "application"
)

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

type httpResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// EncoderResponse encodes the object to the HTTP response.
func EncoderResponse() http.EncodeResponseFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		if rd, ok := v.(http.Redirector); ok {
			url, code := rd.Redirect()
			nethttp.Redirect(w, r, url, code)
			return nil
		}

		// 响应添加 code message
		reply := httpResponse{
			Code:    0,
			Message: "ok",
		}
		codec, _ := http.CodecForRequest(r, "Accept")
		data, err := codec.Marshal(v)
		if err != nil {
			return err
		}
		_ = json.Unmarshal(data, &reply.Data)
		if strings.Contains(r.URL.Path, "/deviceManager/v1/devices") ||
			strings.Contains(r.URL.Path, "/cloudScheduler/v1/streamingservers") {
		} else {
			data, err = codec.Marshal(reply)
		}
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", ContentType(codec.Name()))
		_, err = w.Write(data)
		return err
	}
}

/*
自定义错误映射，主要用于两类场景：
1. 转换框架层自带（比如 jwt）、业务层无法捕获的错误，例如 UNAUTHORIZED 错误
2. 希望使用 http 官方错误状态码，而不是出错时也返回 200
例如：

	func ErrorMapping(err *kratosErrors.Error) *kratosErrors.Error {
		if err.Reason == "UNAUTHORIZED" {
			return kratosErrors.New(400005, "", "用户未认证")
		}
		if err.Code == "430000" {
			err.Metadata = map[string]string{
				"statusCode": "400"
			}
			return err
		}
		return nil
	}
*/
type ErrorMapping func(*errors.Error) *errors.Error

// EncoderError encodes the error to the HTTP response.
func EncoderError(errMapping ErrorMapping) http.EncodeErrorFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
		se := errors.FromError(err)

		var mappedErr *errors.Error
		if errMapping != nil {
			mappedErr = errMapping(se)
		}

		codec, _ := http.CodecForRequest(r, "Accept")
		// 响应添加 code message
		v := &httpResponse{}
		if se.Code == errors.UnknownCode && se.Reason == errors.UnknownReason {
			v.Code = nethttp.StatusInternalServerError
			v.Message = "系统错误"
		} else if mappedErr != nil {
			v.Code = mappedErr.Code
			v.Message = mappedErr.Message
		} else {
			v.Code = se.Code
			v.Message = se.Message
		}
		body, err := codec.Marshal(v)
		if err != nil {
			w.WriteHeader(nethttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", ContentType(codec.Name()))
		if mappedErr != nil && mappedErr.Status.Metadata != nil && mappedErr.Status.Metadata["statusCode"] != "" {
			statusCode := mappedErr.Status.Metadata["statusCode"]
			code, err := strconv.Atoi(statusCode)
			if err != nil {
				w.WriteHeader(nethttp.StatusInternalServerError)
				return
			}
			w.WriteHeader(code)
		} else {
			w.WriteHeader(nethttp.StatusOK)
		}
		_, _ = w.Write(body)
	}
}
