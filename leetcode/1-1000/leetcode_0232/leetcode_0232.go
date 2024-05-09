package leetcode_0232

type MyQueue struct {
	inStack  []int
	outStack []int
	size     int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0, 10),
		outStack: make([]int, 0, 10),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
	this.size++
}

func (this *MyQueue) Pop() int {
	item := this.Peek()
	this.outStack = this.outStack[:len(this.outStack)-1]
	this.size--
	return item
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		panic("queue is empty")
	}
	if len(this.outStack) > 0 {
		// pop outStack
		item := this.outStack[len(this.outStack)-1]
		return item
	}
	// transfer item to outStack
	for i := len(this.inStack) - 1; i >= 0; i-- {
		this.outStack = append(this.outStack, this.inStack[i])
	}
	// clear inStack
	this.inStack = this.inStack[:0]
	// pop outStack
	item := this.outStack[len(this.outStack)-1]
	return item
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}
