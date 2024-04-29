package group

import (
	"fmt"
	"my_local_communitate/pkg/cache/lru"
)

var (
	groups = make(map[string]lru.CacheInterface)
)

func NewGroup(groupName string, cacheType lru.CacheInterface) {
	if _, ok := groups[groupName]; ok {
		return
	}
	groups[groupName] = cacheType
}

func Set(groupName string, key string, value []byte) error {
	if val, ok := groups[groupName]; !ok {
		return fmt.Errorf("can not found group")
	} else {
		val.Set(key, value)
	}
	return nil
}

func Get(groupName string, key string) (value []byte, err error) {
	val, ok := groups[groupName]
	if !ok {
		return nil, fmt.Errorf("can not found group")
	}
	value = val.Get(key)
	if value == nil {
		return nil, fmt.Errorf("can not found value")
	}
	return value, nil
}

func Del(groupName string, key string) {
	val, ok := groups[groupName]
	if ok {
		val.Del(key)
	}
}
