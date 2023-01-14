package leetcode_0641

// https://leetcode.cn/problems/design-circular-deque/

// MyCircularDeque 环形双向队列
type MyCircularDeque struct {
	cap int // 容量
	// [head, tail)
	data []int

	// 类似于  动态数组 的  左闭右开原则
	// 1. 判断队列为空的 条件是: head == tail
	// 2. 判断队列满的   条件是: (tail+1)%cap == head
	head int // 指向队列头部第一个有效数据的索引
	tail int // 指向队列尾部最后一个有效索引的下一个位置
}

// 通过 head tail 指针来 指向 头尾 元素 (底层是一个循环数组)
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap:  k + 1,
		data: make([]int, k+1),
		head: 0,
		tail: 0,
	}
}

func (mq *MyCircularDeque) InsertFront(value int) bool {
	if mq.IsFull() {
		return false
	}
	mq.head = (mq.head - 1 + mq.cap) % mq.cap
	mq.data[mq.head] = value
	return true
}

func (mq *MyCircularDeque) InsertLast(value int) bool {
	if mq.IsFull() {
		return false
	}
	mq.data[mq.tail] = value
	mq.tail = (mq.tail + 1 + mq.cap) % mq.cap
	return true
}

func (mq *MyCircularDeque) DeleteFront() bool {
	if mq.IsEmpty() {
		return false
	}

	mq.head = (mq.head + 1) % mq.cap
	return true

}

func (mq *MyCircularDeque) DeleteLast() bool {
	if mq.IsEmpty() {
		return false
	}

	mq.tail = (mq.tail - 1 + mq.cap) % mq.cap
	return true
}

func (mq *MyCircularDeque) GetFront() int {
	if mq.IsEmpty() {
		return -1
	}
	return mq.data[mq.head]
}

func (mq *MyCircularDeque) GetRear() int {
	if mq.IsEmpty() {
		return -1
	}
	return mq.data[(mq.tail-1+mq.cap)%mq.cap]

}

func (mq *MyCircularDeque) IsEmpty() bool {
	return mq.head == mq.tail
}

func (mq *MyCircularDeque) IsFull() bool {
	return (mq.tail+1)%mq.cap == mq.head
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
