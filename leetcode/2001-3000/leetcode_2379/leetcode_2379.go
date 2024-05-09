package leetcode_2379

// https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
func minimumRecolors(blocks string, k int) int {
	// param check
	if len(blocks) > k {
		// invalid param
		return -1
	}
	var (
		l, r      int
		minWBlock = len(blocks)
		cnt       int
	)
	for r < k {
		if blocks[r] == 'W' {
			cnt++
		}
	}
	minWBlock = min(minWBlock, cnt)
	for r < len(blocks) {
		if blocks[l] == 'W' {
			cnt--
		}
		if blocks[r] == 'W' {
			cnt++
		}
		minWBlock = min(minWBlock, cnt)
		l++
		r++
	}
	return minWBlock
}
