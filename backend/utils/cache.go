package utils

import (
	"sync"
	"time"
)

// NOT　TESTED
// every time put a new value into the map, check the key that expires.
// if the key expires, delete the value from the map.
type Cache struct {
	mem       map[string]any
	volume    int
	minVolume int
	maxVolume int
	timeout   int64 // in seconds
	head      *node
	tail      *node
	sync.RWMutex
}

func NewCache(minVolume int, maxVolume int, timeout int64) *Cache {
	dummy := &node{
		prev:      nil,
		key:       "",
		timestamp: 0,
	}
	return &Cache{
		mem:       make(map[string]any),
		minVolume: minVolume,
		maxVolume: maxVolume,
		volume:    1,
		timeout:   timeout,
		head:      dummy,
		tail:      dummy,
	}
}

type node struct {
	prev      *node
	key       string
	timestamp int64
}

func (cache *Cache) addFirst(key string, value any) {
	// add new node
	cache.head.prev = &node{key: key, prev: nil, timestamp: time.Now().Unix()}
	// move the head pointer
	cache.head = cache.head.prev
	// add to mem
	cache.Lock()
	cache.mem[key] = value
	cache.volume++
	cache.Unlock()
}

func (cache *Cache) removeLast() {
	// remove from mem
	cache.Lock()
	// if _, ok := cache.mem[cache.tail.key]; ok {
	delete(cache.mem, cache.tail.key)
	cache.volume--
	// }
	cache.Unlock()
	// move the tail pointer
	cache.tail = cache.tail.prev
}

func (cache *Cache) trim() {
	// cache.Lock()
	for cache.volume > cache.maxVolume {
		cache.removeLast()
	}
	for cache.volume > cache.minVolume {
		if cache.tail.timestamp+cache.timeout < time.Now().Unix() {
			cache.removeLast()
		}
	}
	// cache.Unlock()
}

func (cache *Cache) Get(key string) (any, bool) {
	cache.RLock()
	value, ok := cache.mem[key]
	cache.RUnlock()
	if ok {
		return value, ok
	}
	return nil, false
}

// only delete keys when put a new node
func (cache *Cache) Put(key string, value any) bool {
	cache.Lock()
	if _, ok := cache.mem[key]; ok {
		cache.mem[key] = value
		cache.Unlock()
		return ok
	}
	cache.Unlock()

	cache.addFirst(key, value)
	cache.trim()
	return true
}

type ListInMem struct {
	sync.RWMutex
}

func NewListInMem() *ListInMem {
	return &ListInMem{
		RWMutex: sync.RWMutex{},
	}
}

func (l *ListInMem) copy() {

	// copy(nil, nil)
}
