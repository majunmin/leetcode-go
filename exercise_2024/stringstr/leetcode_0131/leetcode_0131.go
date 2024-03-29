package leetcode_0131

// https://leetcode.cn/problems/palindrome-partitioning/
func partition(s string) [][]string {
	var result [][]string
	backtrace(&result, s, []string{}, 0, len(s)-1)
	return result
}

func backtrace(result *[][]string, s string, path []string, left, right int) {
	// terminate
	if left > right {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	// 判断前缀是否是 回文子串
	for i := left; i <= right; i++ {
		if isPalindromic(s, left, i) {
			path = append(path, s[left:i+1])
			backtrace(result, s, path, i+1, right)
			path = path[:len(path)-1]
		}
	}
}

func isPalindromic(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
