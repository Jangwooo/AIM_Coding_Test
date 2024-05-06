package utils

import (
	"crypto/rand"
	"fmt"
)

const size = 16

func GenerateNewToken() string {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	key := fmt.Sprintf("%x", b)[:size]
	return key
}
