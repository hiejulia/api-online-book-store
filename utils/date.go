package utils

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

// Timestamp will return a unix timestamp as an unsigned int.
func Timestamp() uint64 {
	ts := time.Now().UTC().Unix()
	return uint64(ts)
}
