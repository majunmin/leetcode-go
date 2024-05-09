package leetcode_0387

// https://leetcode.cn/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	// param check
	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		return 0
	}
	cntArr := make([]int, 26)
	for i := range s {
		cntArr[s[i]-'a']++
	}
	for i := range s {
		if cntArr[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
