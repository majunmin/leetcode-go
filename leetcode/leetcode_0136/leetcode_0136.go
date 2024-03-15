package leetcode_0136

// https://leetcode.cn/problems/single-number/?envType=study-plan-v2&envId=top-100-liked
func singleNumber(nums []int) int {
	// 位运算
	//异或操作
	// param check
	if len(nums) == 0 {
		panic("invalid param, nums.size is 0")
	}
	var result int
	for _, num := range nums {
		result = result ^ num
	}
	return result
}
