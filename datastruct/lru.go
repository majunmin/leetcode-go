// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-20 21:08
package datastruct

type Node struct {
	key   int
	value int

	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	// 维护一个 KV Cache
	cache map[int]*Node

	size int

	// 维护一个双向链表
	firstNode *Node
	lastNode  *Node
}

func Constructor(capacity int) LRUCache {
	firstNode := &Node{}
	lastNode := &Node{}
	firstNode.next = lastNode
	lastNode.prev = firstNode
	return LRUCache{
		capacity:  capacity,
		cache:     make(map[int]*Node),
		firstNode: firstNode,
		lastNode:  lastNode,
	}
}

func (lru *LRUCache) Get(key int) int {
	node, exist := lru.cache[key]
	if !exist {
		return -1
	}
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) moveToHead(node *Node) {
	lru.remove(node)
	node.next = lru.firstNode.next
	lru.firstNode.next.prev = node
	node.prev = lru.firstNode
	lru.firstNode.next = node
}

func (lru *LRUCache) Put(key int, value int) {
	if node, exist := lru.cache[key]; exist {
		node.value = value
		lru.moveToHead(node)
		return
	}
	// 如果 不存在
	node := Node{
		key:   key,
		value: value,
	}
	lru.cache[key] = &node
	lru.moveToHead(&node)
	if lru.size == lru.capacity {
		oldestNode := lru.removeOldest()
		delete(lru.cache, oldestNode.key)
	} else {
		lru.size++
	}
}

func (lru *LRUCache) remove(node *Node) {
	if node == nil || node.prev == nil || node.next == nil {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (lru *LRUCache) removeOldest() *Node {
	// 链表为空
	if lru.firstNode == lru.lastNode {
		return nil
	}
	oldestNode := lru.lastNode.prev
	lru.remove(oldestNode)
	return oldestNode
}
