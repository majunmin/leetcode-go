// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-20 21:08
package lru

type Node struct {
	key   int
	value int

	prev *Node
	next *Node
}

// DoubleLinkedList 定义一个双向链表
type DoubleLinkedList struct {
	head *Node
	tail *Node

	size int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	head := new(Node)
	tail := new(Node)
	head.next = tail
	tail.prev = head

	return &DoubleLinkedList{
		head: head,
		tail: tail,
	}
}

func (dl *DoubleLinkedList) Size() int {
	return dl.size
}

func (dl *DoubleLinkedList) AddFirst(node *Node) {
	dl.head.next.prev = node
	node.next = dl.head.next
	node.prev = dl.head
	dl.head.next = node
	dl.size++
}

func (dl *DoubleLinkedList) AddLast(node *Node) {
	dl.tail.prev.next = node
	node.prev = dl.tail.prev
	node.next = dl.tail
	dl.tail.prev = node
	dl.size++
}

func (dl *DoubleLinkedList) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	// helpGC
	node.next = nil
	node.prev = nil
	dl.size--
}

func (dl *DoubleLinkedList) moveToHead(node *Node) {
	dl.removeNode(node)
	dl.AddFirst(node)
}

// 最老的 在 链表尾部
func (dl *DoubleLinkedList) removeOldest() *Node {
	if dl.size == 0 {
		return nil
	}
	oldestNode := dl.tail.prev
	dl.removeNode(oldestNode)
	return oldestNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node     // 维护一个 KV Cache
	dl       *DoubleLinkedList // 维护一个双向链表

}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		dl:       NewDoubleLinkedList(),
	}
}

func (lru *LRUCache) Get(key int) int {
	node, exist := lru.cache[key]
	if !exist {
		return -1
	}
	lru.dl.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if node, exist := lru.cache[key]; exist {
		node.value = value
		lru.dl.moveToHead(node)
		return
	}

	// 如果 不存在
	node := &Node{
		key:   key,
		value: value,
	}
	lru.dl.AddFirst(node)
	lru.cache[key] = node
	if lru.dl.size > lru.capacity {
		oldestNode := lru.dl.removeOldest()
		delete(lru.cache, oldestNode.key)
	}
}
