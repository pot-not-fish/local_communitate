package communicate

import "my_local_communitate/pkg/crypto/crypto_des"

type Symmetric interface {
	Encrypto(plainData []byte) ([]byte, error)
	Decrypto(cipherData []byte) ([]byte, error)
}

func NewSymmetric(symmetricType string, secretKey string) Symmetric {
	switch symmetricType {
	case "DES":
		return crypto_des.NewCrypto(secretKey)
	}
	return nil
}
