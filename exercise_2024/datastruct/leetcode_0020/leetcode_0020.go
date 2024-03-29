package leetcode_0020

var (
	rightMatch = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
)

// https://leetcode.cn/problems/valid-parentheses/?envType=study-plan-v2&envId=top-100-liked
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if _, exist := rightMatch[s[i]]; !exist {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 || rightMatch[s[i]] != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) > 0
}
