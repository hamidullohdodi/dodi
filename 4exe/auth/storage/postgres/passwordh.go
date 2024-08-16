package postgres

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}