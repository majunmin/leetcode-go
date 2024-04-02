package leetcode_2831

// https://leetcode.cn/problems/find-the-longest-equal-subarray/
func longestEqualSubarray(nums []int, k int) int {
	// param check

	// 1. 构建相同的数字 的position 数组
	num2Pos := make([][]int, len(nums)+1)
	for i, num := range nums {
		num2Pos[num] = append(num2Pos[num], i)
	}
	//
	var ans int
	for _, pos := range num2Pos {
		var l, r int
		for r < len(pos) {
			//
			for (pos[r]-pos[l]+1)-(r-l+1) > k {
				l++
			}
			ans = max(ans, r-l+1)
		}
	}
	return ans
}
