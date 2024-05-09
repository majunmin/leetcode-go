package leetcode_0338

// https://leetcode.cn/problems/counting-bits/
func countBits(n int) []int {

	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//if i&1 == 1 {
		//  // result[i-1] 是偶数, => result[i] = result[(i -1) >> 1] == result[i>>1]
		//	result[i] = result[i-1] + 1
		//} else {
		//  // result[]
		//	result[i] = result[i>>1]
		//}
		result[i] = result[i>>1] + i&1
	}
	return result
}
