package openapi

import "net/http"

type options struct {
	client        client
	header        header
	trimURIPrefix trimURIPrefix
}
type Option interface {
	apply(*options)
}

func newOptions() options {
	return options{
		client:        defaultClient,
		header:        defaultHeader,
		trimURIPrefix: defaultTrimURIPrefix,
	}
}

// Client 自定义http client
type client struct {
	transport *http.Transport
}

var defaultClient = client{
	transport: nil,
}

func (c client) apply(opts *options) {
	opts.client = c
}

func WithTransport(tr *http.Transport) Option {
	return &client{
		transport: tr,
	}
}

// Header 自定义请求头
var alogicHeader = header{
	now:       "x-alogic-now",
	app:       "x-alogic-app",
	ac:        "x-alogic-ac",
	signature: "x-alogic-signature",
	box:       "x-alogic-box",
}

var ctyunaiHeader = header{
	"x-ctyunai-now",
	"x-ctyunai-ak",
	"",
	"x-ctyunai-signature",
	"x-ctyunai-box",
}

var defaultHeader = alogicHeader

type header struct {
	now       string
	app       string
	ac        string
	signature string
	box       string
}

func (h header) apply(opts *options) {
	opts.header = h
}

func WithHeader(now, app, ac, signature, box string) Option {
	return &header{
		now:       now,
		app:       app,
		ac:        ac,
		signature: signature,
		box:       box,
	}
}

// 忽略URI前缀
type trimURIPrefix struct {
	prefix string
}

var defaultTrimURIPrefix = trimURIPrefix{
	prefix: "",
}

func (i trimURIPrefix) apply(opts *options) {
	opts.trimURIPrefix = i
}

func WithTrimURIPrefix(prefix string) Option {
	return &trimURIPrefix{
		prefix: prefix,
	}
}
