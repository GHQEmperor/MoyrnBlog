package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(text string) string {
	obj := sha256.New()
	obj.Write([]byte(text))
	return hex.EncodeToString(obj.Sum(nil))
}
