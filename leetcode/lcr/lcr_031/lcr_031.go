package lcr_031

type node struct {
	key   int
	value int
	onDel func(key, value int)
	next  *node
	prev  *node
}

type doubleLinkedList struct {
	head *node
	tail *node
	size int
}

func newDoubleLinkedList() *doubleLinkedList {
	head, tail := new(node), new(node)
	head.next = tail
	tail.prev = head
	return &doubleLinkedList{
		head: head,
		tail: tail,
	}
}

func (dl *doubleLinkedList) add2Head(item *node) {
	item.next = dl.head.next
	dl.head.next.prev = item
	dl.head.next = item
	item.prev = dl.head
	dl.size++
}

func (dl *doubleLinkedList) move2Head(item *node) {
	dl.remove(item)
	dl.add2Head(item)
}

func (dl *doubleLinkedList) remove(item *node) {
	item.next.prev = item.prev
	item.prev.next = item.next
	// help GC
	item.next = nil
	item.prev = nil
	if item.onDel != nil {
		item.onDel(item.key, item.value)
	}
	dl.size--
}

func (dl *doubleLinkedList) removeOldest() *node {
	x := dl.tail.prev
	dl.remove(x)
	return x
}

func (dl *doubleLinkedList) len() int {
	return dl.size
}

// https://leetcode.cn/problems/OrIXps/
type LRUCache struct {
	cache map[int]*node
	dl    *doubleLinkedList
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*node),
		dl:    newDoubleLinkedList(),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, exist := this.cache[key]; !exist {
		return -1
	}
	item := this.cache[key]
	this.dl.move2Head(item)
	return item.value
}

func (this *LRUCache) Put(key int, value int) {
	if item, exist := this.cache[key]; exist {
		// update
		item.value = value
		this.dl.move2Head(item)
		return
	}
	if this.cap >= this.Len() {
		removed := this.dl.removeOldest()
		delete(this.cache, removed.key)
	}
	newNode := &node{
		key:   key,
		value: value,
		onDel: nil,
	}
	this.cache[key] = newNode
	this.dl.add2Head(newNode)
}
func (this *LRUCache) Len() int {
	return this.dl.len()
}
