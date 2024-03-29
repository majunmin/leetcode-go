package leetcode_1456

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
func maxVowels(s string, k int) int {
	vowelSet := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	// 定长滑动窗口
	//param Check
	var result int
	var cnt int
	k = min(k, len(s))
	for i := 0; i < k; i++ {
		if vowelSet[s[i]] {
			cnt++
		}
	}
	result = cnt
	// 滑动窗口
	for i := k; i < len(s); i++ {
		if vowelSet[s[i]] == true {
			cnt++
		}
		if vowelSet[s[i-k]] {
			cnt--
		}
		result = max(result, cnt)
	}
	return result
}
