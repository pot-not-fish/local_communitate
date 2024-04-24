package communicate

import "my_local_communitate/pkg/crypto/crypto_md5"

type Check interface {
	IsValid(data []byte, signature string) bool
}

func NewCheck(checkType string) Check {
	switch checkType {
	case "md5":
		return crypto_md5.NewHashMD5()
	}
	return nil
}
