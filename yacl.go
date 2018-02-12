package yacl

import (
	"time"
)

type cache struct {
	store      map[string]item
	items      []item
	expiration time.Duration
}

type item struct {
	key        string
	value      interface{}
	inserted   time.Time
	expiration time.Duration
}

func NewCache() *cache {
	s := make(map[string]item)
	c = &cache{s, []item{}, time.Duration(time.Second * 5)}

	go func() {
		t := time.Tick(1 * time.Second)
		for _ = range t {
			now := time.Now()
			for _, i := range c.items {
				if now.After(i.inserted.Add(i.expiration)) {
					c.items = c.items[1:]
					delete(c.store, i.key)
				} else {
					break
				}
			}
		}
	}()
	return c
}

func (c *cache) Insert(key string, value interface{}, args ...interface{}) {
	var expiration time.Duration
	if len(args) == 0 {
		expiration = c.expiration
	} else {
		expiration = args[0].(time.Duration)
	}
	c.items = append(c.items, item{key, value, time.Now(), expiration})
	c.store[key] = item{key, value, time.Now(), expiration}
}

func (c *cache) Get(key string) interface{} {
	if i, ok := c.store[key]; ok {
		return i.value
	}
	return nil
}
