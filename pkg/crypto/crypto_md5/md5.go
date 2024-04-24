package crypto_md5

import (
	"crypto/md5"
	"fmt"
)

type hashMD5 struct{}

func NewHashMD5() *hashMD5 {
	return new(hashMD5)
}

func (h *hashMD5) IsValid(data []byte, signature string) bool {
	hash := md5.Sum(data)
	hex := fmt.Sprintf("%x", hash)
	return hex == signature
}
