package leetcode_3206

// https://leetcode.cn/problems/alternating-groups-i/
func numberOfAlternatingGroups(colors []int) int {
	var (
		result int
		n      = len(colors)
	)
	for i := 0; i < len(colors); i++ {
		if colors[i] == colors[(i-1+n)%n] || colors[i] == colors[(i+1+n)%n] {
			continue
		}
		result++
	}
	return result
}
