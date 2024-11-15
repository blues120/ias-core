package signature

import "reflect"

var DefaultHeader = header{
	now:       "x-alogic-now",
	app:       "x-alogic-app",
	signature: "x-alogic-signature",
	box:       "x-alogic-box",
}

var CtyunaiHeader = header{
	now:       "x-ctyunai-now",
	app:       "x-ctyunai-ak",
	signature: "x-ctyunai-signature",
	box:       "x-ctyunai-box",
}

var OpenApiHeader = header{
	now:       "x-ctyunai-now",
	app:       "x-ctyunai-ak",
	signature: "x-ctyunai-signature",
	orgId:     "x-ctyunai-org",
}

type options struct {
	H header
}

func NewOptions() options {
	return options{
		H: DefaultHeader,
	}
}

type Option interface {
	Apply(*options)
}

type header struct {
	now       string
	app       string
	signature string
	box       string
	orgId     string
}

func (h header) Apply(opts *options) {
	opts.H = h
}

func (h header) Keys() []string {
	var keys []string
	hVal := reflect.ValueOf(h)
	for i := 0; i < hVal.NumField(); i++ {
		if key := hVal.Field(i).String(); key != "" {
			keys = append(keys, key)
		}
	}
	return keys
}

func (h header) GetSignatureKey() string {
	return h.signature
}

// WithHeader 自定义请求头
func WithHeader(now, app, signature, box string) Option {
	return &header{
		now:       now,
		app:       app,
		signature: signature,
		box:       box,
	}
}
