package leetcode_0705

import (
	"hash"
	"hash/fnv"
	"strconv"
)

// 1. hash 函数设计， 减少hash冲突
// 2. 扩容 负载因子
// 3. hash 冲突的解决

type entry struct {
	key  int
	next *entry
}
type MyHashSet struct {
	table []*entry
	hash  hash.Hash64
}

func Constructor() MyHashSet {
	return MyHashSet{
		table: make([]*entry, 1024),
		hash:  fnv.New64(),
	}
}

func (this *MyHashSet) Add(key int) {
	idx := this.tableIdx(key)
	if this.table[idx] == nil {
		// insert head
		this.table[idx] = &entry{key: key}
	} else {
		if this.table[idx].key == key {
			return
		}
		pre := this.table[idx]
		// 发生hash冲突
		for ; pre.next != nil; pre = pre.next {
			if pre.next.key == key {
				return
			}
		}
		pre.next = &entry{key: key}
	}
}

func (this *MyHashSet) tableIdx(key int) int {
	this.hash.Write([]byte(strconv.Itoa(key)))
	hashVal := this.hash.Sum64()
	this.hash.Reset()
	idx := int(hashVal) & (len(this.table) - 1)
	return idx
}

func (this *MyHashSet) Remove(key int) {
	idx := this.tableIdx(key)
	e := this.table[idx]
	if e == nil {
		return
	}
	if e.key == key {
		this.table[idx] = e.next
	} else {
		pre := e
		for ; pre.next != nil; pre = pre.next {
			if pre.next.key == key {
				pre.next = pre.next.next
			}
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	idx := this.tableIdx(key)
	if this.table[idx] == nil {
		return false
	}
	for e := this.table[idx]; e != nil; e = e.next {
		if e.key == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
