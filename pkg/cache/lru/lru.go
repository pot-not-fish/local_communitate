package lru

import (
	"container/list"
	"sync"
)

const (
	maxCacheCount = 100
)

type CacheLRU struct {
	mu           sync.Mutex
	currentCount int
	ll           *list.List
	cache        map[string]*list.Element
}

type valueLRU struct {
	key   string
	value []byte
}

type CacheInterface interface {
	Get(key string) (value []byte)
	Set(key string, value []byte)
}

func NewCache() *CacheLRU {
	return &CacheLRU{
		ll:    list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (c *CacheLRU) Get(key string) (value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.cache[key]; ok {
		value = val.Value.(*valueLRU).value
		c.ll.MoveToFront(val)
	}
	return value
}

func (c *CacheLRU) Set(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.cache[key]; ok {
		c.ll.MoveToFront(val)
		val.Value.(*valueLRU).value = value
		return
	}
	if c.currentCount+1 > maxCacheCount {
		c.del()
	}
	head := c.ll.PushFront(&valueLRU{key: key, value: value})
	c.cache[key] = head
}

func (c *CacheLRU) del() {
	value := c.ll.Back()
	delete(c.cache, value.Value.(*valueLRU).key)
	c.ll.Remove(value)
}
