package leetcode_0225

type MyStack struct {
	curQueue  []int
	helpQueue []int
}

func Constructor() MyStack {
	return MyStack{
		curQueue:  make([]int, 0, 10),
		helpQueue: make([]int, 0, 10),
	}
}

func (this *MyStack) Push(x int) {
	this.helpQueue = append(this.helpQueue, x)
	for i := 0; i < len(this.curQueue); i++ {
		this.helpQueue = append(this.helpQueue, this.curQueue[i])
	}
	// clear curQueue
	this.curQueue = this.curQueue[:0]
	this.curQueue, this.helpQueue = this.helpQueue, this.curQueue
}

func (this *MyStack) Pop() int {
	item := this.Top()
	this.curQueue = this.curQueue[1:]
	return item
}

func (this *MyStack) Top() int {
	if len(this.curQueue) == 0 {
		panic("myStack is empty")
	}
	item := this.curQueue[0]
	return item
}

func (this *MyStack) Empty() bool {
	return len(this.curQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
