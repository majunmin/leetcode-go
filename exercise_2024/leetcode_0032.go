package exercise_2024

// https://leetcode.cn/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	//
	var maxLen int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		if s[i] == ')' {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				stack = stack[:len(stack)-1]
				// 更新result
				if len(stack) > 0 {
					maxLen = max(maxLen, i-stack[len(stack)-1])
				}
			}
		}
	}
	return maxLen
}
