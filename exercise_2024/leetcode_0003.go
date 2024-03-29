package exercise_2024

func lengthOfLongestSubstring(s string) int {
	// param check
	if len(s) < 2 {
		return len(s)
	}
	var result int
	cnt := make(map[byte]int)
	var l, r int
	for r < len(s) {
		if cnt[s[r]] == 0 {
			cnt[s[r]]++
			r++
			continue
		}
		result = max(result, r-l)
		// shrink window
		for l < r {
			cnt[s[l]]--
			if cnt[s[l]] == 1 {
				l++
				break
			}
			l++
		}
	}
	result = max(result, r-l)
	return result
}
