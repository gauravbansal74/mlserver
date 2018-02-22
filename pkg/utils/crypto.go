package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	iteration = 25000
	keyLen    = 512
)

// GetHash - Get hash using sha1
func GetHash(source string, salt string) (string, error) {
	if source == "" {
		return "", fmt.Errorf("source can't be null or empty")
	}
	if salt == "" {
		return "", fmt.Errorf("salt can't be null or empty")
	}
	hash := pbkdf2.Key([]byte(source), []byte(salt), iteration, keyLen, sha1.New)
	return hex.EncodeToString(hash), nil
}
