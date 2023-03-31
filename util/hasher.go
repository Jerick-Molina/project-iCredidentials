package util

import (
	"crypto/sha256"
	"fmt"
)

func Hasher(passwd string) string {

	sum := sha256.Sum256([]byte(passwd))

	return fmt.Sprintf("%x", sum)
}
