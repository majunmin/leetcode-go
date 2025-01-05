package leetcode_3208

// https://leetcode.cn/problems/alternating-groups-ii/
func numberOfAlternatingGroups(colors []int, k int) int {
	var (
		result int
		cnt    int
		n      = len(colors)
	)

	for i := 0; i < n*2; i++ {
		if i > 0 && colors[i%n] != colors[(i-1+n)%n] {
			cnt = 0
		}
		cnt++
		if i >= n && cnt >= k {
			result++
		}
	}
	return result
}
