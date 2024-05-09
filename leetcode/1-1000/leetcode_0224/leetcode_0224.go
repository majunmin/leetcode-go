package leetcode_0224

import (
	"strings"
)

// https://leetcode.cn/problems/basic-calculator/
func calculate(s string) int {
	// param check 校验表达式合法性
	opStack := newStack[byte]()
	numStack := newStack[int]()
	s = strings.ReplaceAll(s, " ", "")
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			var num int
			for i < len(s) && isDigit(s[i]) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i--
			numStack.push(num)
		} else {
			switch s[i] {
			case '(':
				opStack.push(s[i])
			case ')':
				for !opStack.isEmpty() && opStack.top() != '(' {
					calc(opStack, numStack)
				}
				opStack.pop() // 弹出左括号
			default: // 运算符
				// -1-2-3 ==> 0-1-2-3
				// 1-2-(-3+1) ==> 1-2-(0-3+1)
				// 1-2-(+3+1) ==> 1-2-(0+3+1)
				if i == 0 || !opStack.isEmpty() && s[i-1] == '(' { // 对于 单操作数的运算符的优化
					numStack.push(0)
				}
				for !opStack.isEmpty() && priority(opStack.top()) >= priority(s[i]) {
					calc(opStack, numStack)
				}
				// 将栈中可以算的操作符 都算了,(todo: 优先级)
				opStack.push(s[i])
			}
		}
	}
	for !opStack.isEmpty() {
		calc(opStack, numStack)
	}
	return numStack.top()

}

func calc(opStack *stack[byte], numStack *stack[int]) {
	if numStack.len() < 2 {
		return
	}
	b, a := numStack.pop(), numStack.pop() // 弹出两个运算数
	op := opStack.pop()                    // 弹出运算符
	//fmt.Printf("cal %d %d %d \n", a, op, b)
	switch op {
	case '+':
		numStack.push(a + b)
	case '-':
		numStack.push(a - b)
	case '*':
		numStack.push(a * b)
	case '/':
		numStack.push(a / b)
	}
}

type stack[T int | byte] struct {
	items []T
}

func newStack[T int | byte]() *stack[T] {
	c := stack[T]{}
	c.items = make([]T, 0)

	return &c
}

func (s *stack[T]) pop() T {
	x := s.items[s.len()-1]
	s.items = s.items[:s.len()-1]
	return x
}

func (s *stack[T]) push(t T) {
	s.items = append(s.items, t)
}

func (s *stack[T]) top() T {
	if s.isEmpty() {
		return 0
	}
	return s.items[s.len()-1]
}

func (s *stack[T]) len() int {
	return len(s.items)
}

func (s *stack[T]) isEmpty() bool {
	return s.len() == 0
}

func isDigit(item byte) bool {
	return item >= '0' && item <= '9'
}

func priority(op byte) int {
	switch op {
	case '-', '+':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func calculate2(s string) int {
	opsStack := newStack[int]()
	sign := 1
	var result int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
		case '+':
			sign = opsStack.top()
		case '-':
			sign = -opsStack.top()
		case '(':
			opsStack.push(sign)
		case ')':
			opsStack.pop()
		default:
			// isDigit
			var num int
			for i < len(s) && isDigit(s[i]) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
		}
	}
	return result
}
