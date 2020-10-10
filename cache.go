package main

import "sync"

// Cache :
type Cache struct {
	sync.RWMutex
	data map[string]interface{}
}

// NewCache :
func NewCache() *Cache {
	cache := &Cache{
		data:    make(map[string]interface{}, 1000),
		RWMutex: sync.RWMutex{},
	}
	return cache
}

// Add :
func (c *Cache) Add(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

// Get :
func (c *Cache) Get(key string) (value interface{}, doesExist bool) {
	c.RLock()
	defer c.RUnlock()
	value, doesExist = c.data[key]
	return
}
