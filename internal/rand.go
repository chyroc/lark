package internal

import (
	"math/rand"
	"strings"
	"time"
)

func RandString(count int) string {
	s := strings.Builder{}

	for i := 0; i < count; i++ {
		s.WriteByte(randStringSource[rand.Intn(len(randStringSource))])
	}
	return s.String()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randStringSource = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
