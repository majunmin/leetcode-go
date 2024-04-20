package leetcode_0707

type node struct {
	val  int
	prev *node
	next *node
}

// https://leetcode.cn/problems/design-linked-list/
type MyLinkedList struct {
	head *node
	tail *node
	size int
}

func Constructor() MyLinkedList {
	head, tail := new(node), new(node)
	head.next = tail
	tail.prev = head
	return MyLinkedList{
		head: head,
		tail: tail,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if !this.checkIndex(index) {
		return -1
	}
	//
	e := this.head
	for i := 0; i <= index; i++ {
		e = e.next
	}
	return e.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &node{
		val: val,
	}
	this.insertAfter(this.head, newNode)
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &node{val: val}
	this.insertAfter(this.tail.prev, newNode)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 {
		// param invalid
		return
	}
	newNode := &node{val: val}
	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	this.insertAfter(prev, newNode)
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if !this.checkIndex(index) {
		return
	}
	evict := this.head
	for i := 0; i <= index; i++ {
		evict = evict.next
	}
	evict.next.prev = evict.prev
	evict.prev.next = evict.next
	// help GC
	evict.next = nil
	evict.prev = nil
	this.size--
}

func (this *MyLinkedList) checkIndex(index int) bool {
	return index < this.size && index >= 0
}

func (this *MyLinkedList) insertAfter(prev *node, e *node) {
	e.next = prev.next
	prev.next.prev = e
	e.prev = prev
	prev.next = e
	this.size++
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
