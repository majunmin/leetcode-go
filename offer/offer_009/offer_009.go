package offer_009

type stack []int

func NewStack(initSize int) stack {
	return make([]int, 0, initSize)
}

func (s *stack) Push(item int) {
	*s = append(*s, item)
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

func (s stack) isEmpty() bool {
	return len(s) == 0

}

// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/?plan=lcof&plan_progress=z2ptdrc
type CQueue struct {
	inStack  stack
	outStack stack
}

func Constructor() CQueue {
	return CQueue{
		inStack:  NewStack(10),
		outStack: NewStack(10),
	}
}

func (cq *CQueue) AppendTail(value int) {
	cq.inStack.Push(value)
}

func (cq *CQueue) DeleteHead() int {
	if !cq.outStack.isEmpty() {
		return cq.outStack.Pop()
	}

	if cq.inStack.isEmpty() {
		return -1
	}

	for !cq.inStack.isEmpty() {
		cq.outStack.Push(cq.inStack.Pop())
	}
	return cq.outStack.Pop()

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
