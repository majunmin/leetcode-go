package leetcode_100245

// https://leetcode.cn/contest/weekly-contest-390/problems/maximum-length-substring-with-two-occurrences/
func maximumLengthSubstring(s string) int {

	cnts := make([]int, 26)
	var l, r int
	var result int
	for r < len(s) {
		r++
		cnts[s[r-1]-'a']++
		if cnts[s[r-1]-'a'] <= 2 {
			result = max(result, r-l)
			continue
		}
		// shrink window
		for l < r {
			l++
			cnts[s[l-1]-'a']--
			if cnts[s[l-1]-'a'] == 2 {
				break
			}
		}
	}
	return result
}
