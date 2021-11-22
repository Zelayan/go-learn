// lru
package lru

import (
	"container/list"
)

type Cache struct {
	// maxBytes is max use memory
	maxBytes  int64
	// nbytes is use memory at now
	nbytes    int64
	// ll is 双向链表
	ll        *list.List
	// cache is dict
	cache     map[string]*list.Element
	// OnEvicted 删除某条记录的回调函数，可以为nil
	OnEvicted func(key string, value Value)
}

type entry struct {
	key string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (Value Value, ok bool) {
	if element, ok := c.cache[key]; ok{
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		return kv.value,true
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes = c.nbytes
	}
}

