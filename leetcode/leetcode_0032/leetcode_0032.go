package leetcode_0032

// https://leetcode.cn/problems/longest-valid-parentheses/?envType=study-plan-v2&envId=top-100-liked
func longestValidParentheses(s string) int {
	return dpSolution(s)
}

func stackSolution(s string) int {
	// param check
	if len(s) < 2 {
		return 0
	}
	var result int
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		item := s[i]
		if item == '(' || stack[len(stack)-1] == -1 || s[stack[len(stack)-1]] != '(' {
			stack = append(stack, i)
			continue
		}
		// item == ')'
		stack = stack[:len(stack)-1]
		top := stack[len(stack)-1]
		result = max(result, i-top)
	}
	return result
}

func dpSolution(s string) int {
	// param check
	if len(s) < 2 {
		return 0
	}
	//
	result := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			// '()'保底 dp[i] = 2
			dp[i] = 2
			// 如果存在 )()  的场景, 那么 dp[i] = dp[i-2] + 2
			//         (()
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
			// '(()())'保底 dp[i] = dp[i-1] + 2
			//  '())' 或者 ')())'的场景不会出现
			dp[i] = dp[i-1] + 2
			if i-dp[i-1] >= 2 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		result = max(result, dp[i])
	}
	return result
}
