package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(value *string) string {
	if value == nil {
		return ""
	}
	sum := sha256.Sum256([]byte(*value))
	return hex.EncodeToString(sum[:])
}
