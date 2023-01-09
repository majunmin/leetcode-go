package leetcode

type MinStack struct {
	stack  []int
	minVal []int
}

func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		minVal: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	if len(ms.stack) == 0 || val <= ms.minVal[len(ms.minVal)-1] {
		ms.minVal = append(ms.minVal, val)
	} else {
		ms.minVal = append(ms.minVal, ms.minVal[len(ms.minVal)-1])
	}
	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minVal = ms.minVal[:len(ms.minVal)-1]
}

func (ms *MinStack) Top() int {
	// return -1 if stack is empty
	if len(ms.stack) == 0 {
		return -1
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	// return -1 if stack is empty
	if len(ms.minVal) == 0 {
		return -1
	}
	return ms.minVal[len(ms.minVal)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
