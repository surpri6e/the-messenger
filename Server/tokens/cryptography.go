package token

import (
	"crypto/sha256"
	"fmt"
)

func Crypto(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hex := h.Sum(nil)
	hexString := fmt.Sprintf("%x", hex)
	return hexString
}
