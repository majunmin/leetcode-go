// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-24 23:26
package leetcode_0155

import "math"

// 1.利用一个辅助栈,  保存  最小值
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: []int{math.MaxInt64},
	}

}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, minVal(val, this.minStack[len(this.minStack)-1]))
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
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func minVal(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
