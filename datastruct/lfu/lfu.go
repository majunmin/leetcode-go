// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-21 23:45
package lfu

type Node struct {
	key   int
	value int
	// 访问频次 计数
	frequency int

	prev *Node
	next *Node
}

type DoubleLinkList struct {
	// 双向链表 维护访问顺序
	firstNode *Node
	lastNode  *Node
}

func NewDoubleLinkList(frequency int) *DoubleLinkList {
	firstNode := Node{}
	lastNode := Node{}
	firstNode.next = &lastNode
	lastNode.prev = &firstNode
	firstNode.frequency = frequency
	lastNode.frequency = frequency
	return &DoubleLinkList{
		firstNode: &firstNode,
		lastNode:  &lastNode,
	}
}

func (l *DoubleLinkList) IsEmpty() bool {
	return l.firstNode.next == l.lastNode
}

func (l *DoubleLinkList) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	// helpGC
	node.prev = nil
	node.next = nil
}

//func (l *DoubleLinkList) addToHead(node *Node) {
//	l.firstNode.next.prev = node
//	node.next = l.firstNode.next
//	node.prev = l.firstNode
//	l.firstNode.next = node
//}

type LFUCache struct {
	capacity int
	size     int

	cache         map[int]*Node
	frequencyList []*DoubleLinkList // 维护访问频次的顺序

}

func Constructor(capacity int) LFUCache {
	frequencyList := []*DoubleLinkList{NewDoubleLinkList(1)}

	return LFUCache{
		capacity:      capacity,
		size:          0,
		cache:         make(map[int]*Node, capacity),
		frequencyList: frequencyList,
	}
}

func (lfu *LFUCache) Get(key int) int {
	node, exist := lfu.cache[key]
	if !exist {
		return -1
	}
	// 升级
	lfu.remove(node)
	node.frequency++
	if len(lfu.frequencyList) < node.frequency {
		//需要扩容
		lfu.frequencyList = append(lfu.frequencyList, NewDoubleLinkList(node.frequency))
	}
	if lfu.size == lfu.capacity {
		lfu.removeOldest()
	}

	lfu.addToHead(node)
	//lfu.cache[key] = node
	return node.value
}

func (lfu *LFUCache) Put(key int, value int) {
	// condition
	if lfu.capacity == 0 {
		return
	}
	node, exist := lfu.cache[key]
	if exist {
		node.value = value
		lfu.remove(node)
	} else {
		node = &Node{
			key:   key,
			value: value,
		}
	}

	node.frequency++
	if len(lfu.frequencyList) < node.frequency {
		//需要扩容
		lfu.frequencyList = append(lfu.frequencyList, NewDoubleLinkList(node.frequency))
	}
	if lfu.size == lfu.capacity {
		lfu.removeOldest()
	}

	lfu.addToHead(node)
	lfu.cache[key] = node
}

func (lfu *LFUCache) addToHead(node *Node) {
	if node.frequency > len(lfu.frequencyList) {
		lfu.frequencyList = append(lfu.frequencyList, NewDoubleLinkList(node.frequency))
	}
	l := lfu.frequencyList[node.frequency-1]
	l.firstNode.next.prev = node
	node.next = l.firstNode.next
	node.prev = l.firstNode
	l.firstNode.next = node
	lfu.cache[node.key] = node
	lfu.size++
}

func (lfu *LFUCache) remove(node *Node) {
	dList := lfu.frequencyList[node.frequency-1]
	dList.remove(node)
	// del
	delete(lfu.cache, node.key)
	lfu.size--
}

func (lfu *LFUCache) removeOldest() {
	for _, dList := range lfu.frequencyList {
		if dList.IsEmpty() {
			continue
		}
		lfu.remove(dList.lastNode.prev)
		break
	}
}
