package leetcode

import (
	"container/heap"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	var (
		n        = len(s)
		numStack = newStack[int]()
		opStack  = newStack[byte]()
	)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var u int
			for i < n && isDigit(s[i]) {
				u = u*10 + int(s[i]-'0')
				i++
			}
			numStack.push(u)
			i--
		case '(':
			opStack.push('(')
		case ')':
			// 计算到 左边的左括号
			for opStack.len() > 0 && opStack.top() != '(' {
				calc(numStack, opStack)
			}
			opStack.pop() // pop '('
		default:
			// op
			if i == 0 || (s[i-1] == '(' || s[i-1] == '+' || s[i-1] == '-') {
				numStack.push(0)
			}
			// todo: 优先级处理
			for opStack.len() > 0 && opStack.top() != '(' && priority(opStack.top()) > priority(s[i]) {
				calc(numStack, opStack)
			}
			opStack.push(s[i])
		}
	}
	for opStack.len() > 0 {
		calc(numStack, opStack)
	}
	return numStack.pop()
}

func calc(numStack *stack[int], opStack *stack[byte]) {
	n2 := numStack.pop()
	n1 := numStack.pop()
	op := opStack.pop()
	switch op {
	case '+':
		numStack.push(n1 + n2)
	case '-':
		numStack.push(n1 - n2)
	case '*':
		numStack.push(n1 * n2)
	case '/':
		numStack.push(n1 / n2)
	}
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}

type stack[T int | byte] struct {
	items []T
}

func newStack[T int | byte]() *stack[T] {
	return &stack[T]{
		items: make([]T, 0),
	}
}

func (s *stack[T]) push(x T) {
	s.items = append(s.items, x)
}

func (s *stack[T]) pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *stack[T]) top() T {
	if s.len() == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *stack[T]) len() int {
	return len(s.items)
}

func priority(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 1
	}
}

func myAtoi(s string) int {
	var (
		hasSign bool
		sign    = true //表示正数
		res     int
	)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if hasSign && s[i] == '+' || s[i] == '-' {
			break
		}
		if s[i] == '+' {
			hasSign = true
		} else if s[i] == '-' {
			hasSign = true
			sign = false
		} else if isDigit(s[i]) {
			res = res*10 + int(s[i]-'0')
		} else {
			break
		}
	}
	if !sign {
		return -res
	}
	return res
}

type MedianFinder struct {
	leftHp  *maxHeap
	rightHp *minHeap
}

/** initialize your data structure here. */
func Constructor2() MedianFinder {
	return MedianFinder{
		leftHp:  newMaxHeap(),
		rightHp: newMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.leftHp.Len() == this.rightHp.Len() {
		heap.Push(this.rightHp, num)
		heap.Push(this.leftHp, heap.Pop(this.rightHp).(int))
	} else {
		heap.Push(this.leftHp, num)
		heap.Push(this.rightHp, heap.Pop(this.leftHp).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.leftHp.Len() != this.rightHp.Len() {
		// fmt.Println(*this.leftHp)
		return float64((*this.leftHp)[0])
	} else {
		return float64((*this.leftHp)[0]+(*this.rightHp)[0]) / 2.0
	}
}

type minHeap []int

func newMinHeap() *minHeap {
	hp := new(minHeap)
	heap.Init(hp)
	return hp
}
func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

type maxHeap []int

func newMaxHeap() *maxHeap {
	hp := new(maxHeap)
	heap.Init(hp)
	return hp
}
func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Push(x any) {
	//TODO implement me
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
