package leetcode_100255

import "sort"

// https://leetcode.cn/contest/weekly-contest-389/problems/minimum-deletions-to-make-string-k-special/
func minimumDeletions(word string, k int) int {
	cnts := make([]int, 26)
	for i := range word {
		cnts[word[i]-'a']++
	}
	sort.Ints(cnts)
	var maxSave int
	// process
	for i, base := range cnts {
		var sum int
		for _, c := range cnts[i:] {
			sum += min(c, base+k)
		}
		maxSave = max(maxSave, sum)
	}
	return len(word) - maxSave
}
