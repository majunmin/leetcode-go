package leetcode_0146

import (
	"container/list"
	"fmt"
)

func OnEvictDefault(key, value int) {
	fmt.Println(key, ":", value)
}

type valueWrap struct {
	key   int
	value int
}

type LRUCache struct {
	cap     int
	ll      *list.List
	cache   map[int]*list.Element
	onEvict func(key, value int)
}

func Constructor(capacity int) LRUCache {
	// param check
	if capacity <= 0 {
		panic("invalid capacity")
	}
	return LRUCache{
		cap:     capacity,
		ll:      list.New(),
		cache:   make(map[int]*list.Element),
		onEvict: nil,
	}
}

func (lru *LRUCache) Get(key int) int {
	if _, exist := lru.cache[key]; !exist {
		return -1
	}
	// exist
	// move to list.head
	lru.ll.MoveToFront(lru.cache[key])
	return lru.cache[key].Value.(*valueWrap).value
}

func (lru *LRUCache) Put(key int, value int) {
	if e, exist := lru.cache[key]; exist {
		e.Value.(*valueWrap).value = value
		lru.ll.MoveToFront(e)
		return
	}
	// if full
	if lru.cap <= len(lru.cache) {
		lru.removeOldest()
	}
	// add
	lru.addToHead(key, value)
}

func (lru *LRUCache) removeOldest() {
	val := lru.ll.Back()
	lru.ll.Remove(val)
	lru.cap--
	valWrap := val.Value.(*valueWrap)
	delete(lru.cache, valWrap.key)
	if lru.onEvict != nil {
		lru.onEvict(valWrap.key, valWrap.value)
	}
}

func (lru *LRUCache) addToHead(key int, value int) {
	val := valueWrap{
		key:   key,
		value: value,
	}
	ee := lru.ll.PushFront(&val)
	lru.cache[key] = ee
	//lru.cap++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
