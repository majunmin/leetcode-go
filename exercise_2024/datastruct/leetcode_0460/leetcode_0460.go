package leetcode_0460

type node struct {
	key   int
	val   interface{}
	prev  *node
	next  *node
	level int
}

func newNode(k int, v interface{}) *node {
	return &node{
		key:   k,
		val:   v,
		level: 0,
	}
}

type DLinkedList struct {
	head *node
	tail *node
	size int
}

func NewDLinkedList() *DLinkedList {
	head, tail := new(node), new(node)
	head.next = tail
	tail.prev = head
	return &DLinkedList{
		head: head,
		tail: tail,
	}
}

func (dl *DLinkedList) Len() int {
	return dl.size
}

func (dl *DLinkedList) remove(node *node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	// help gc
	node.next = nil
	node.prev = nil
	dl.size--
}

func (dl *DLinkedList) removeOldest() *node {
	oldNode := dl.tail.prev
	dl.remove(oldNode)
	return oldNode
}

func (dl *DLinkedList) add2Head(node *node) {
	// node should not be nil
	node.next = dl.head.next
	dl.head.next.prev = node
	dl.head.next = node
	node.prev = dl.head
	dl.size++
}

type LFUCache struct {
	cache map[int]*node
	// 使用切片来维护 访问次数
	levelList []*DLinkedList
	cap       int
	size      int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:     make(map[int]*node),
		levelList: make([]*DLinkedList, 0),
		size:      0,
		cap:       capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, exist := this.cache[key]; !exist {
		//
		return -1
	}
	node := this.cache[key]
	this.levelList[node.level].remove(node)
	node.level++
	if len(this.levelList) == node.level {
		this.levelList = append(this.levelList, NewDLinkedList())
	}
	this.levelList[node.level].add2Head(node)
	return node.val.(int)
}

func (this *LFUCache) Put(key int, value int) {
	if node, exist := this.cache[key]; exist {
		this.levelList[node.level].remove(node)
		node.val = value
		node.level++
		if len(this.levelList) == node.level {
			this.levelList = append(this.levelList, NewDLinkedList())
		}
		this.levelList[node.level].add2Head(node)
		return
	}
	// insert into new level
	nNode := newNode(key, value)
	if len(this.levelList) == 0 {
		this.levelList = append(this.levelList, NewDLinkedList())
	}
	this.cache[key] = nNode
	this.levelList[0].add2Head(nNode)
	if this.cap == this.size {
		for i := 0; i < len(this.levelList); i++ {
			if this.levelList[i].Len() > 0 {
				n := this.levelList[i].removeOldest()
				delete(this.cache, n.key)
				this.size--
				break
			}
		}
	}
	this.size++
}
