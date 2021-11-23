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

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
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
		c.nbytes = c.nbytes - int64(len(kv.key)) + int64(kv.value.Len())

		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add
func (c *Cache) Add(key string, value Value) {
	// if the key exists，update value and nbytes
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		// if not exists, push it to front and join the cache
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele

		c.nbytes += int64(len(key)) + int64(value.Len())
	}
    // if nbytes more than maxbytes,then remove oldest
	if c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}

}


