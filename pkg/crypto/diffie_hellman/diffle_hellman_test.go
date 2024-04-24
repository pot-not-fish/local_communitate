package diffie_hellman

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCrypto(t *testing.T) {
	data := int64(25864)

	cytA := NewCrypto()
	cytB := NewCrypto()

	cytA.gAB = QuickPower(cytB.gR, cytA.r)
	cytB.gAB = QuickPower(cytA.gR, cytB.r)

	cipher := cytA.Encrypto(data)

	plain := cytB.Decrypto(cipher)
	fmt.Println(data, cipher, plain)
	assert.Equal(t, plain, data)
}
