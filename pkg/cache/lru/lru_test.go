package lru

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLRU(t *testing.T) {
	m := NewCache()
	for i := 0; i < 100; i++ {
		m.Set(fmt.Sprintf("%d", i), []byte("a"))
	}
	m.Set("0", []byte("a"))
	m.Set("1", []byte("b"))
	m.Set("100", []byte("a"))
	m.Set("101", []byte("a"))

	assert.Equal(t, m.Get("0"), []byte("a"))
	assert.Equal(t, m.Get("1"), []byte("b"))
	assert.Equal(t, m.ll.Back().Value.(*valueLRU).key, "2")
	assert.Equal(t, m.Get("-1"), nil)
}
