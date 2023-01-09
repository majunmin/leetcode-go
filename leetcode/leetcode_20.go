package leetcode

var (
	m = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
)

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[b] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}
