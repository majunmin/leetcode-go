package leetcode_0438

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	needs := [26]int{}
	for i := range p {
		needs[p[i]-'a']++
	}
	//
	window := [26]int{}
	var l, r int
	var result []int
	for r < len(s) {
		window[s[r]-'a']++
		r++
		// shrink window
		for r-l == len(p) {
			if window == needs {
				result = append(result, l)
			}
			window[s[l]-'a']--
			l++
		}
	}
	return result
}
