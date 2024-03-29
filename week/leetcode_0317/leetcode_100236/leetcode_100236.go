package leetcode_100236

// https://leetcode.cn/contest/weekly-contest-389/problems/count-substrings-starting-and-ending-with-given-character/
func countSubstrings(s string, c byte) int64 {
	//
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	result := 1
	for cnt > 1 {
		result *= cnt
		cnt--
	}
	return int64(result)
}
