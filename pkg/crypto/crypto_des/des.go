package crypto_des

import (
	"github.com/forgoer/openssl"
)

type crypto struct {
	secretKey []byte
}

func NewCrypto(secretKey string) *crypto {
	cyt := new(crypto)
	cyt.secretKey = []byte(secretKey)
	if len(cyt.secretKey) > 8 {
		cyt.secretKey = cyt.secretKey[:8]
	} else {
		for len(cyt.secretKey) != 8 {
			cyt.secretKey = append(cyt.secretKey, '0')
		}
	}
	return cyt
}

func (c *crypto) Encrypto(plainData []byte) ([]byte, error) {
	return openssl.DesCBCEncrypt(plainData, c.secretKey, c.secretKey, openssl.PKCS7_PADDING)
}

func (c *crypto) Decrypto(cipherData []byte) ([]byte, error) {
	return openssl.DesCBCDecrypt(cipherData, c.secretKey, c.secretKey, openssl.PKCS7_PADDING)
}
