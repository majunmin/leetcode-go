package leetcode_0146

type Node struct {
	key  int
	val  interface{}
	prev *Node
	next *Node
}

type onEvict func(k int, v interface{}) error

func newNode() *Node {
	return &Node{}
}

func newKVNode(k int, v interface{}) *Node {
	return &Node{key: k, val: v}
}

type DoubleLinkList struct {
	head    *Node
	tail    *Node
	size    int
	onEvict onEvict
}

func NewDoubleLinkList() *DoubleLinkList {
	head, tail := newNode(), newNode()
	head.next = tail
	tail.prev = head
	return &DoubleLinkList{
		head: head,
		tail: tail,
	}
}

func (dl *DoubleLinkList) setOnEvict(onEvict onEvict) {
	dl.onEvict = onEvict
}

func (dl *DoubleLinkList) len() int {
	return dl.size
}

func (dl *DoubleLinkList) isEmpty() bool {
	return dl.size == 0
}

func (dl *DoubleLinkList) addToHead(node *Node) {
	dl.head.next.prev = node
	node.next = dl.head.next
	dl.head.next = node
	node.prev = dl.head
	dl.size++
}

func (dl *DoubleLinkList) Remove(node *Node) {
	dl.remove(node)
	if dl.onEvict != nil {
		dl.onEvict(node.key, node.val)
	}
}

func (dl *DoubleLinkList) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	// help gc
	node.next = nil
	node.prev = nil
	dl.size--
}

func (dl *DoubleLinkList) removeOldest() *Node {
	oldNode := dl.tail.prev
	dl.remove(oldNode)
	return oldNode
}

func (dl *DoubleLinkList) moveToHead(node *Node) {
	dl.remove(node)
	dl.addToHead(node)
}

// 利用map 存储 kv 映射关系
// 利用 list存储访问时间顺序
type LRUCache struct {
	cache map[int]*Node
	list  *DoubleLinkList
	cap   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*Node),
		list:  NewDoubleLinkList(),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, exist := this.cache[key]; !exist {
		return -1
	}
	node := this.cache[key]
	this.list.moveToHead(node)
	return node.val.(int)
}

func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.cache[key]; !exist {
		node.val = value
		this.list.moveToHead(node)
		return
	}
	// insert new node
	if this.list.len() == this.cap {
		// expire oldest
		node := this.list.removeOldest()
		delete(this.cache, node.key)
	}
	node := newKVNode(key, value)
	this.list.addToHead(node)
	this.cache[key] = node
}
