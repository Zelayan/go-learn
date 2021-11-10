package main

import "container/list"

func main() {

}

type entry struct {
	Key, Value int
}

type LRUCache struct {
	cap int
	cache map[int]*list.Element
	lst *list.List
}

func Constructir(capacity int) LRUCache {
	return LRUCache{
		capacity,
		map[int]*list.Element{},
		list.New(),
	}
}

func (c *LRUCache)Get(key int) int {
	e := c.cache[key]
	if e == nil {
		return -1
	}
	// 刷新缓存时间，将元素提到最前面
	c.lst.MoveToFront(e)
	return  e.Value.(entry).Value
}

func (c *LRUCache) Put(key, value int) {
	// can find it
	if e := c.cache[key]; e != nil {
		e.Value = entry{key, value}
		// reflush cache time
		c.lst.MoveToFront(e)
		return
	}

	c.cache[key] = c.lst.PushFront(entry{key, value})
	// 删除最久未使用元素
	if len(c.cache) > c.cap {
		delete(c.cache, c.lst.Remove(c.lst.Back()).(entry).Key)
	}

}

