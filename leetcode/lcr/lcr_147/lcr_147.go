package lcr_147

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
type MinStack struct {
	minStack []int
	stack    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack: []int{},
		stack:    []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 || x < this.minStack[len(this.minStack)-1] {
		this.stack = append(this.stack, x)
		return
	}
	this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])

}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	item := this.stack[len(this.stack)-1]
	return item
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return -1
	}
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
