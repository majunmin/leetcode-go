package leetcode_0706

import (
	"hash"
	"hash/fnv"
	"strconv"
)

type entry struct {
	key  int
	val  int
	next *entry
}

type MyHashMap struct {
	table []*entry
	size  int
	hash  hash.Hash32
}

func Constructor() MyHashMap {
	return MyHashMap{
		table: make([]*entry, 1024),
		hash:  fnv.New32(),
		size:  0,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.tableIndex(key)
	if this.table[idx] == nil {
		// 1. not exist
		e := entry{key: key, val: value}
		this.table[idx] = &e
		return
	}
	// 头插或者尾插,尾插需要记录尾指针
	// 这里采用头插法
	for e := this.table[idx]; e != nil; e = e.next {
		if e.key == key {
			e.val = value
			return
		}
	}
	insertData := &entry{key: key, val: value}
	insertData.next = this.table[idx]
	this.table[idx] = insertData
	this.size++
}

func (this *MyHashMap) Get(key int) int {
	idx := this.tableIndex(key)
	if this.table[idx] == nil {
		// 1. not exist
		return -1
	}
	for e := this.table[idx]; e != nil; e = e.next {
		if e.key == key {
			return e.val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := this.tableIndex(key)
	if this.table[idx] == nil {
		// 1. not exist
		return
	}
	e := this.table[idx]
	if e.key == key {
		this.table[idx] = e.next
		this.size--
	} else {
		for pre := e; pre.next != nil; pre = pre.next {
			if pre.next.key == key {
				pre.next = pre.next.next
				this.size--
				return
			}
		}
	}
}

func (this *MyHashMap) tableIndex(key int) int {
	this.hash.Write([]byte(strconv.Itoa(key)))
	idx := this.hash.Sum32()
	this.hash.Reset()
	return int(idx) & (1<<10 - 1)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
