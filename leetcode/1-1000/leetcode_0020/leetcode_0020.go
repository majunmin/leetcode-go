// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-19 22:43
package leetcode_0020

//https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {

	// 如果 len(s) 是奇数 一定返回false
	if len(s) == 0 {
		return false
	}
	parenthesesMatch := make(map[byte]byte)
	parenthesesMatch[')'] = '('
	parenthesesMatch[']'] = '['
	parenthesesMatch['}'] = '{'
	// 一般这种括号匹配的问题 可以用栈来解决
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if c, exist := parenthesesMatch[s[i]]; exist {
			// 遇到右括号 栈为空 || 栈顶元素不匹配， 快速失败
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			// 匹配则 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 遇到左括号入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
