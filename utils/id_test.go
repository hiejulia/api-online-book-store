package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestID(t *testing.T) {
	id := ID()
	assert.NotNil(t, id)
}
