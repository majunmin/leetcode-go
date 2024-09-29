package leetcode_3129

// https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-i/
func numberOfStableArrays(zero int, one int, limit int) int {
	// 连续 0 or 1 的个数不能超过 limit +1.
	//
	var (
		path = make([]int, 0)
	)

}

// 用 zero个0 和 one个1 构成的稳定数组方案数.
// 其中第 i+j个位置要填k.
func dfs(zero, one int, limit int, k int, path []int) int {
	// terminate
	if i == len(path) {
		return 1
	}
	var ret int
	if zero > 0 {

	}
}
