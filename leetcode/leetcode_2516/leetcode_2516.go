package leetcode_2516

// https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/
func takeCharacters(s string, k int) int {
	// param check
	if len(s) == 0 && k == 0 || len(s) < 3*k {
		return -1
	}
	var (
		needs = map[byte]int{
			'a': 0,
			'b': 0,
			'c': 0,
		}
		l, r     int
		maxWSize int
	)
	for i := 0; i < len(s); i++ {
		needs[s[i]]++
	}
	// 窗口内的 字符数 需要满足  num(b) <= needs[i],
	// 求窗口的最长的
	for b, cnt := range needs {
		if cnt < k {
			return -1
		}
		needs[b] -= k
	}
	for ; r < len(s); r++ {
		needs[s[r]]--
		for !valid(needs) {
			needs[s[l]]++
			l++
		}
		maxWSize = max(maxWSize, r-l+1)
	}
	return len(s) - maxWSize
}

func valid(needs map[byte]int) bool {
	for _, cnt := range needs {
		if cnt < 0 {
			return false
		}
	}
	return true
}
