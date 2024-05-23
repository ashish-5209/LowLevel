package lruCache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	ll       *list.List
}

type Entry struct {
	key, value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	if elem, found := c.cache[key]; found {
		c.ll.MoveToFront(elem)
		return elem.Value.(*Entry).value, true
	}
	return 0, false
}

func (c *LRUCache) Put(key, value int) {
	if elem, found := c.cache[key]; found {
		c.ll.MoveToFront(elem)
		elem.Value.(*Entry).value = value
		return
	}
	if c.ll.Len() >= c.capacity {
		oldest := c.ll.Back()
		if oldest != nil {
			c.ll.Remove(oldest)
			delete(c.cache, oldest.Value.(*Entry).key)
		}
	}

	newEntry := &Entry{key, value}
	elem := c.ll.PushFront(newEntry)
	c.cache[key] = elem
}

func App() {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1, true
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns 0, false
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns 0, false
	fmt.Println(cache.Get(3)) // returns 3, true
	fmt.Println(cache.Get(4)) // returns 4, true
}
