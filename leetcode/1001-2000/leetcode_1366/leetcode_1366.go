package leetcode_1366

import (
	"cmp"
	"maps"
	"slices"
)

// https://leetcode.cn/problems/rank-teams-by-votes/?envType=daily-question&envId=2024-12-30
func rankTeams(votes []string) string {
	cnts := make(map[byte][]int)

	for i := range votes[0] {
		cnts[votes[0][i]] = make([]int, len(votes[0]))
	}

	for _, vote := range votes {
		for i := range vote {
			cnts[vote[i]][i]++
		}
	}
	//
	result := slices.SortedFunc(maps.Keys(cnts), func(a, b byte) int {
		return cmp.Or(slices.Compare(cnts[b], cnts[a]), cmp.Compare(a, b))
	})
	return string(result)
}
