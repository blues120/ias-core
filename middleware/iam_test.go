//go:build skip
// +build skip

package middleware

import (
	"fmt"
	"testing"

	"github.com/blues120/ias-core/middleware/signature"
)

func TestOption(t *testing.T) {
	options := signature.NewOptions()
	signature.OpenApiHeader.Apply(&options)
	fmt.Println(options.H.Keys())
}
