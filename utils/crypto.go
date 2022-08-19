package utils

import (
	"crypto/sha256"
	"fmt"
)

// data -> hex
func Hash(data interface{}) string {
	str := fmt.Sprintf("%v", data)
	hash := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", hash)
}
