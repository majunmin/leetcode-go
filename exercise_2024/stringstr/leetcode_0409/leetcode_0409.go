package leetcode_0409

// https://leetcode.cn/problems/longest-palindrome/
func longestPalindrome(s string) int {
	// 利用map 统计每个字符出现的次数.
	byteCnt := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, exist := byteCnt[s[i]]; !exist {
			byteCnt[s[i]] = 1
		}
		byteCnt[s[i]] += 1
	}

	// 计算最长回文子串
	var flag bool //是否存在奇数个字符
	var result int
	for _, cnt := range byteCnt {
		if cnt&1 == 0 {
			result += cnt
		} else {
			result += cnt - 1
		}
	}
	if flag {
		result += 1
	}
	return result
}
