package lru

import "container/list"

const (
	maxCacheCount = 100
)

type cacheLRU struct {
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

func NewCache() *cacheLRU {
	return &cacheLRU{
		ll:    list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (c *cacheLRU) Get(key string) (value []byte) {
	if val, ok := c.cache[key]; ok {
		value = val.Value.(*valueLRU).value
		c.ll.MoveToFront(val)
	}
	return value
}

func (c *cacheLRU) Set(key string, value []byte) {
	if val, ok := c.cache[key]; ok {
		c.ll.MoveToFront(val)
		val.Value.(*valueLRU).value = value
	}
	if c.currentCount+1 > maxCacheCount {
		c.del()
	}
	head := c.ll.PushFront(&valueLRU{key: key, value: value})
	c.cache[key] = head
}

func (c *cacheLRU) del() {
	value := c.ll.Back()
	delete(c.cache, value.Value.(*valueLRU).key)
	c.ll.Remove(value)
}
