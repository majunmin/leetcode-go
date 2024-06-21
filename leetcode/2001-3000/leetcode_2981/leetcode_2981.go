package leetcode_2981

// https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-i/
func maximumLength(s string) int {
	n := len(s)
	if n < 3 {
		return -1
	}
	// 遍历特殊字符的长度 i.
	for i := n - 2; i > 0; i-- {
		// 记录每个特殊字符的长度.
		cnts := make(map[string]int)
		for j := 0; j <= n-i; j++ {
			if !isTeshu(s[j : j+i]) {
				continue
			}
			cnts[s[j:j+i]]++
			if cnts[s[j:j+i]] >= 3 {
				return i
			}
		}
	}
	return -1
}

func isTeshu(s string) bool {
	p := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != p {
			return false
		}
	}
	return true
}
