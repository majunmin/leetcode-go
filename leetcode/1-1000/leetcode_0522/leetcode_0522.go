package leetcode_0522

// https://leetcode.cn/problems/longest-uncommon-subsequence-ii/
func findLUSlength(strs []string) int {
	var (
		n      = len(strs)
		result int
	)
	for i := 0; i < n; i++ {
		flag := true // strs[i] 不是其他字符串的子序列
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if subSequence(strs[i], strs[j]) {
				flag = false
				break
			}
		}
		if flag {
			result = max(result, len(strs[i]))
		}
	}
	return result
}

// subStr是否是 str的子序列
func subSequence(sub string, str string) bool {
	m, n := len(sub), len(str)
	if m > n {
		return false
	}
	var matchIdx int
	for i := range str {
		if matchIdx == m-1 {
			return true
		}
		if str[i] == sub[matchIdx] {
			matchIdx++
		}
	}
	return matchIdx == m-1
}

//https://leetcode.cn/problems/longest-uncommon-subsequence-ii/solutions/1627846/by-ac_oier-vuez/
