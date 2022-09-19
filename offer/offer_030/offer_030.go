package offer_030

//https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/?plan=lcof&plan_progress=z2ptdrc

type stack []int

func NewStack(initSize int) stack {
	return make([]int, 0, initSize)
}

func (s *stack) Push(val int) {
	*s = append(*s, val)
}

func (s *stack) Pop() int {
	sLen := len(*s)
	if sLen == 0 {
		return -1
	}
	item := (*s)[sLen-1]
	*s = (*s)[:sLen-1]
	return item
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Top() int {
	if len(*s) == 0 {
		return -1
	}
	return (*s)[len(*s)-1]
}

type MinStack struct {
	norStack stack
	minStack stack // 维护最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		norStack: NewStack(10),
		minStack: NewStack(10),
	}
}

func (ms *MinStack) Push(x int) {
	ms.norStack.Push(x)
	if ms.minStack.isEmpty() || x <= ms.minStack.Top() {
		ms.minStack.Push(x)
		return
	}
	ms.minStack.Push(ms.minStack.Top())
}

func (ms *MinStack) Pop() {
	ms.norStack.Pop()
	ms.minStack.Pop()
}

func (ms *MinStack) Top() int {
	return ms.norStack.Top()
}

func (ms *MinStack) Min() int {
	if ms.minStack.isEmpty() {
		return -1
	}
	return ms.minStack.Top()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
