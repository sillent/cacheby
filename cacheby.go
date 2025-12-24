package cacheby

import "sync"

type CacheBy[K, K2 comparable, V any] struct {
	rw   sync.RWMutex
	data map[K]map[K2]*V
}

func NewCacheBy[K, K2 comparable, V any]() *CacheBy[K, K2, V] {
	return &CacheBy[K, K2, V]{
		rw:   sync.RWMutex{},
		data: make(map[K]map[K2]*V),
	}
}

func (c *CacheBy[K, K2, V]) Store(byKeyName K, key K2, val V) {
	c.rw.Lock()
	defer c.rw.Unlock()
	if c.data[byKeyName] == nil {
		c.data[byKeyName] = make(map[K2]*V)
	}
	c.data[byKeyName][key] = &val
}

func (c *CacheBy[K, K2, V]) Load(byKeyName K, key K2) *V {
	c.rw.RLock()
	defer c.rw.RUnlock()
	byKey, exist := c.data[byKeyName]
	if !exist {
		return nil
	}
	val, exist := byKey[key]
	if !exist {
		return nil
	}
	return val
}

func (c *CacheBy[K, K2, V]) ListBy(key K) map[K2]*V {
	c.rw.RLock()
	defer c.rw.RUnlock()
	kk, exist := c.data[key]
	if exist {
		return kk
	}
	return nil
}
