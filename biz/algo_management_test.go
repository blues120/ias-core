//go:build skip
// +build skip

package biz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateField(t *testing.T) {
	sourceAlgo := &Algorithm{
		Version:          "2.0",
		AlgoGroupName:    "g1",
		AlgoGroupVersion: "2.0",
	}
	destAlgo := &Algorithm{
		Version:          "1.0",
		AlgoGroupName:    "g1",
		AlgoGroupVersion: "1.0",
	}
	UpdateField(sourceAlgo, destAlgo, "AlgoGroupVersion")
	assert.Equal(t, "2.0", destAlgo.AlgoGroupVersion, "group version changed")
}
