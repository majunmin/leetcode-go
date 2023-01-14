package datastruct

type Node struct {
	prev *Node
	next *Node

	// 访问频次  lfu
	frequency int // 维护访问频次

	key   string
	value string
}

type DoubleLinkList struct {
	head *Node
	tail *Node

	size int
}

func NewDoubleLinkList() *DoubleLinkList {
	head, tail := new(Node), new(Node)
	head.next = tail
	tail.prev = head

	return &DoubleLinkList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (dl *DoubleLinkList) AddFirst(node *Node) {
	if node == nil {
		return
	}
	dl.head.next.prev = node
	node.next = dl.head.next
	dl.head.next = node
	node.prev = dl.head
	dl.size++
}

func (dl *DoubleLinkList) AddLast(node *Node) {
	if node == nil {
		return
	}

	dl.tail.prev.next = node
	node.prev = dl.tail.prev.next
	dl.tail.prev = node
	node.next = dl.tail
	dl.size++
}

func (dl *DoubleLinkList) moveToHead(node *Node) {
	dl.removeNode(node)
	dl.AddFirst(node)
}

func (dl *DoubleLinkList) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	// helpGC
	node.next = nil
	node.prev = nil
	dl.size--
}

func (dl *DoubleLinkList) removeOldest() *Node {
	if dl.size == 0 {
		return nil
	}
	node := dl.tail.prev
	dl.removeNode(node)
	return node
}

type LRUCache struct {
	dl       *DoubleLinkList
	cache    map[string]*Node
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	dl := NewDoubleLinkList()
	cache := make(map[string]*Node, capacity)
	return &LRUCache{
		dl:    dl,
		cache: cache,
	}
}

func (c *LRUCache) Put(key, value string) {
	if node, exist := c.cache[key]; exist {
		node.value = value
		c.dl.moveToHead(node)
	}
	node := &Node{
		key:   key,
		value: value,
	}
	c.dl.AddFirst(node)
	if c.Size() == c.capacity {
		removeNode := c.dl.removeOldest()
		delete(c.cache, removeNode.key)
	}
}

func (c *LRUCache) Get(key string) string {
	if node, exist := c.cache[key]; !exist {
		return ""
	} else {
		c.dl.moveToHead(node)
		return node.value
	}
}

func (c LRUCache) Size() int {
	return c.dl.size
}
