package leetcode_0316

// https://leetcode.cn/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}

	stack := make([]byte, 0)
	visit := [26]bool{}
	for i := range s {
		ch := s[i]
		if !visit[ch-'1'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				visit[last] = false
			}
			stack = append(stack, ch)
			visit[ch-'a'] = true
		}
		left[ch-'a']--
	}

	return string(stack)
}
