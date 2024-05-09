// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-11 21:23
package leetcode_1036

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// https://leetcode-cn.com/problems/escape-a-large-maze/
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(source) != 2 || len(target) != 2 ||
		source[0] < 0 || source[0] >= 10000_000 || source[1] < 0 || source[1] >= 10000_000 ||
		target[0] < 0 || target[0] >= 10000_000 || target[1] < 0 || target[1] >= 10000_000 {
		return false
	}

	visitedSet := make(map[int64]struct{}, 10^6*10^6)
	blockedSet := make(map[int64]struct{}, len(blocked))
	for _, blockItem := range blocked {
		blockedSet[int64(blockItem[0])*10^6+int64(blockItem[1])] = struct{}{}
	}
	return backTrace(source[0], source[1], target[0], target[1], blockedSet, visitedSet)
}

func backTrace(s0, s1, t0, t1 int, blockedSet map[int64]struct{}, visitedSet map[int64]struct{}) bool {
	// terminate
	if s0 < 0 || s0 >= 10000_000 || s1 < 0 || s1 >= 10000_000 ||
		t0 < 0 || t0 >= 10000_000 || t1 < 0 || t1 >= 10000_000 {
		return false
	}

	item := int64(s0)*1_000_000 + int64(s1)
	// 剪枝
	if _, exist := blockedSet[item]; exist {
		return false
	}

	if _, exist := visitedSet[item]; exist {
		return false
	}
	visitedSet[item] = struct{}{}

	if s0 == t0 && s1 == t1 {
		return true
	}

	// for choice in choiceList
	for _, dir := range directions {
		news0, news1 := s0+dir[0], s1+dir[1]
		if backTrace(news0, news1, t0, t1, blockedSet, visitedSet) {
			return true
		}
	}
	delete(visitedSet, item)
	return false
}
