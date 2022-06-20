package cache

import (
	"time"
)

type Cache struct {
	dict  map[string]string
	times map[string]time.Time
}

func NewCache() Cache {
	return Cache{
		dict:  make(map[string]string),
		times: make(map[string]time.Time),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	deadline, isDeadlineFound := c.times[key]
	if deadline.Before(time.Now()) && isDeadlineFound {
		delete(c.dict, key)
		delete(c.times, key)
	}
	value, isFound := c.dict[key]
	return value, isFound
}

func (c *Cache) Put(key, value string) {
	c.dict[key] = value
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.dict))
	for k := range c.dict {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.dict[key] = value
	c.times[key] = deadline
}
