package exercise_2024

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/?envType=study-plan-v2&envId=top-100-liked
func findAnagrams(s string, p string) []int {
	// param check
	if len(p) == 0 || len(s) < len(p) {
		return nil
	}
	pcnt := [26]int{}
	for i := 0; i < len(p); i++ {
		pcnt[p[i]-'a']++
	}
	cnt := [26]int{}
	var l, r int
	result := make([]int, 0, len(s)-len(p)+1)
	for r < len(s) {
		cnt[s[r]-'a']++
		r++
		if r-l == len(p) {
			if cnt == pcnt {
				result = append(result, l)
			}
			cnt[s[l]-'a']--
			l++
		}
	}
	return result
}
