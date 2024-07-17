package leetcode_0256

// https://leetcode.ca/all/256.html#google_vignette
func minCost(costs [][]int) int {
	// param check

	var r, g, b int
	// r, g, b
	for _, cost := range costs {
		_r := min(g, b) + cost[0]
		_g := min(r, b) + cost[1]
		_b := min(r, g) + cost[2]
		r, g, b = _r, _g, _b
	}
	return min(r, g, b)
}
