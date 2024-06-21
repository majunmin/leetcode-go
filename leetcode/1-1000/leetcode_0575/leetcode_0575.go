package leetcode_0575

// https://leetcode.cn/problems/distribute-candies/
func distributeCandies(candyType []int) int {
	var set = make(map[int]struct{}, len(candyType))
	for _, typ := range candyType {
		set[typ] = struct{}{}
	}
	return min(len(set), len(candyType)/2)
}
