package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenerateSHA256Hash generates a SHA-256 hash for the given password
func GenerateSHA256Hash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
