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

type Crypto struct {
	r   int64
	gR  int64
	gAB int64
}

func NewCrypto() *Crypto {
	cyt := new(Crypto)
	cyt.r = GetRandom()
	cyt.gR = QuickPower(g, cyt.r)
	return cyt
}

func (c *Crypto) KeyGeneration(keyGeneration func(c Crypto) int64) {
	c.gAB = keyGeneration(*c)
}

func (c *Crypto) Get_R() int64 {
	return c.r
}

func (c *Crypto) Get_GR() int64 {
	return c.gR
}

func (c *Crypto) Get_gAB() int64 {
	return c.gAB
}

func (c *Crypto) Encrypto(data int64) int64 {
	return (data * c.gAB) % p
}

func (c *Crypto) Decrypto(data int64) int64 {
	return (data * QuickPower(c.gAB, p-2)) % p
}

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
