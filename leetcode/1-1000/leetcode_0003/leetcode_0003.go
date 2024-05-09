package leetcode_0003

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&envId=top-100-liked
func lengthOfLongestSubstring(s string) int {

	// param check
	//skip
	charMap := make(map[byte]int)
	var result int
	//var windowSize int
	var l, r int
	for ; r < len(s); r++ {
		charMap[s[r]]++
		if charMap[s[r]] == 1 {
			//windowSize++
			result = max(result, r-l+1)
			continue
		}
		// charMap[s[r]]  == 2
		// 缩短左边界
		for l < r {
			lItem := s[l]
			l++
			charMap[lItem]--
			if charMap[lItem] == 1 {
				//windowSize--
				break
			}
			// charMap[s[l]] == 1
			// skip
		}

	}
	return result
}

// 题解1:https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/228576/longest-substring-without-repeating-characters-b-2/
