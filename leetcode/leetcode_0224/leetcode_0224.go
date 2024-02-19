package leetcode_0224

// https://leetcode.cn/problems/basic-calculator/
func calculate(s string) int {
	// param check 校验表达式合法性
	opStack := newStack[byte]()
	numStack := newStack[int]()
	for i := range s {
		item := s[i]
		if item == ' ' {
			continue
		}
		if item >= '0' && item <= '9' {
			numStack.push(int(item - '0'))
		}
		switch item {
		case '(':
			opStack.push(item)
		case ')':
		case '-':
		case '+':
			if opStack.len() > 0 && opStack.top() != '(' {
				calc(opStack, numStack)
			}
			opStack.push(item)
		}
	}
}

func calc(opStack *stack[byte], numStack *stack[int]) int {
	for opStack.len() > 0 && opStack.top() != '(' {
		op := opStack.pop()
		switch op {
		case '+':
			num1 := numStack.pop()
			num2 := numStack.pop()
			numStack.push(num1 + num2)
		case '-':
			num1 := numStack.pop()
		}
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

func isDigit(item byte) bool {
	return item >= '0' && item <= '9'
}
