//go:build skip
// +build skip

package middleware

import (
	"fmt"
	"testing"

	"gitlab.ctyuncdn.cn/ias/ias-core/middleware/signature"
)

func TestOption(t *testing.T) {
	options := signature.NewOptions()
	signature.OpenApiHeader.Apply(&options)
	fmt.Println(options.H.Keys())
}
