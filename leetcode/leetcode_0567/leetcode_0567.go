package leetcode_0567

// https://leetcode.cn/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	// param check
	if len(s1) > len(s2) {
		return false
	}
	// needs : s1 的 所有字符数
	// 如果 两个 字符串排列相同， 那么有  slices.Equal(needs1, needs2)
	needs := [26]int{}
	for i := range s1 {
		needs[s1[i]]++
	}
	window := [26]int{}
	var l, r int
	for r < len(s2) {
		window[s2[r]]++
		r++

		// shrink window
		for r-l == len(s1) {
			if needs == window {
				return true
			}
			//
			window[s2[l]]--
			l++
		}
	}
	return false
}
