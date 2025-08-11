package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	created_at time.Time
	val []byte
}

type Cache struct {
	mu sync.RWMutex
	cacheMap map[string]cacheEntry
	//Add(string, []byte)
	//Get(string) ([]byte,bool)
}

func (c Cache) Add(key string, value []byte){
	c.mu.Lock()
    defer c.mu.Unlock()
	ce := cacheEntry{created_at: time.Now(), val: value}
	c.cacheMap[key] = ce
}

func (c Cache) Get(key string) ([]byte, bool){
	c.mu.RLock()
    defer c.mu.RUnlock()
	ret, ok := c.cacheMap[key]
	return ret.val, ok
}

func (c Cache) readLoop(interval time.Duration){
	for key,val := range c.cacheMap{
		if time.Since(val.created_at) > interval{
			//c.cacheMap[key] = cacheEntry{}
			delete(c.cacheMap, key)
		}
	}
	
}

//func NewCache(interval time.Duration)