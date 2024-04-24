package communicate

import (
	"my_local_communitate/pkg/crypto/diffie_hellman"
)

type Asymmetric interface {
	Encrypto(data int64) int64
	Decrypto(data int64) int64
}

func NewAsymmetric(asymmetricType string) Asymmetric {
	switch asymmetricType {
	case "diffie_hellman":
		return diffie_hellman.NewCrypto()
	}
	return nil
}
