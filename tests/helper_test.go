package tests

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ptrString(s string) *string {
	return &s
}

func randInt64() int64 {
	return rand.Int63()
}
