package leetcode_0316

// https://leetcode.cn/problems/remove-duplicate-letters/
// 官方题解写的比较好.
// 1. 如果字母是升序排列的， name结果的字典序就是最小的
//
// 2. 如果当前字符已经在栈中了， 则不能将当前字符加入栈中
// 3. 弹出栈顶元素时,如果后面已经没有该栈顶元素的字符了, 那么栈顶元素不能弹出.
func removeDuplicateLetters(s string) string {
	// param check
	// 记录每个字母 最后出现的索引.(用于判断该字母时候还会在后面出现)
	lastIndex := make([]int, 26)
	for i := range s {
		lastIndex[s[i]-'a'] = i
	}

	// process
	stack := make([]byte, 0)
	visit := make([]bool, 26)
	for i := range s {
		ch := s[i]
		// 在栈中已经存在， 跳过
		if visit[int(ch-'a')] {
			continue
		}

		for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastIndex[stack[len(stack)-1]-'a'] {
			// pop
			peekItem := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			visit[peekItem-'a'] = false
		}

		visit[ch-'a'] = true
		stack = append(stack, ch)
	}
	return string(stack)
}

// 第一版
// 1. 用一个hash表计数:  用于 查询该 character是否在后面还会出现
func solution1(s string) string {
	// param check
	// 记录每个字母出现的次数
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	// process
	stack := make([]byte, 0)
	visit := make(map[byte]bool)
	for i := range s {
		ch := s[i]
		// 在栈中已经存在， 跳过
		if visit[ch] {
			m[ch]--
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch {
			peekItem := stack[len(stack)-1]
			if m[peekItem] == 0 {
				break
			}
			// pop
			stack = stack[:len(stack)-1]
			visit[peekItem] = false
		}

		m[ch]--
		visit[ch] = true
		stack = append(stack, ch)
	}
	return string(stack)
}
