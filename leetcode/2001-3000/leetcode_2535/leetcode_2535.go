package leetcode_2535

// https://leetcode.cn/problems/difference-between-element-sum-and-digit-sum-of-an-array/
func differenceOfSum(nums []int) int {
	var result int
	for _, num := range nums {
		var digitSum int
		for num > 0 {
			digitSum += num % 10
			digitSum /= 10
		}
		result += num - digitSum
	}
	return result
}
