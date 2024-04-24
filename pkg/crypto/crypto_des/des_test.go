package crypto_des

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDes(t *testing.T) {
	cyt1 := NewCrypto("228170")
	assert.Equal(t, cyt1.secretKey, []byte("22817000"))

	cyt2 := NewCrypto("22817000000000000000")
	assert.Equal(t, cyt2.secretKey, []byte("22817000"))

	plaintext := "hello world"
	cipherData, err := cyt1.Encrypto([]byte(plaintext))
	assert.Equal(t, err, nil)

	plainData, err := cyt1.Decrypto(cipherData)
	assert.Equal(t, err, nil)
	fmt.Println(string(plainData))
	assert.Equal(t, string(plainData), plaintext)
}
