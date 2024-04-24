package diffie_hellman

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	g = 3
	p = 167772161
)

type crypto struct {
	r   int64
	gR  int64
	gAB int64
}

func NewCrypto() *crypto {
	cyt := new(crypto)
	cyt.r = GetRandom()
	cyt.gR = QuickPower(g, cyt.r)
	cyt.KeyGenerator()
	return cyt
}

func (c *crypto) Encrypto(data int64) int64 {
	return (data * c.gAB) % p
}

func (c *crypto) Decrypto(data int64) int64 {
	return (data * QuickPower(c.gAB, p-2)) % p
}

func (c *crypto) KeyGenerator() {}

func GetRandom() (random int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := md5.Sum([]byte(fmt.Sprintf("%d%d", r.Intn(100000), time.Now().UnixNano())))
	hex := fmt.Sprintf("%x", hash)
	random, _ = strconv.ParseInt(hex, 16, 64)
	return random % p
}

func QuickPower(base int64, pow int64) (res int64) {
	res = 1
	factor := base
	for pow != 0 {
		if pow&1 == 1 {
			res *= factor
		}
		factor *= factor
		factor %= p
		res %= p
		pow >>= 1
	}
	return res
}
