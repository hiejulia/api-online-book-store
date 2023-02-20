package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimestampSuccess(t *testing.T) {
	ts := Timestamp()
	assert.NotNil(t, ts)
}
