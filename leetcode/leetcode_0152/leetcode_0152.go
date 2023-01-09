package leetcode_0152

// https://leetcode.cn/problems/maximum-product-subarray/description/
func maxProduct(nums []int) int {
	return dpSolution2(nums)

}

// 根据 动态规划的 无后效性,可以  得出 优化空间的解法, 以当前位置结尾的 乘积最大值 只跟 前一个位置的极值有关
func dpSolution2(nums []int) int {

	// param check
	if len(nums) == 0 {
		return 0
	}

	//
	fMax, fMin := nums[0], nums[0]
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		tmpFMax, tmpFMin := fMax, fMin
		fMax = maxInt(nums[i], maxInt(nums[i]*tmpFMax, nums[i]*tmpFMin))
		fMin = minInt(nums[i], minInt(nums[i]*tmpFMax, nums[i]*tmpFMin))
		maxVal = maxInt(maxVal, fMax)
	}
	return maxVal
}

// 动态规划解法
func dpSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dpMin := make([]int, len(nums))
	dpMax := make([]int, len(nums))
	dpMin[0], dpMax[0] = nums[0], nums[0]
	maxVal := nums[0]

	for i := 1; i < len(nums); i++ {
		dpMax[i] = maxInt(nums[i], maxInt(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i]))
		dpMin[i] = minInt(nums[i], minInt(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i]))
		maxVal = maxInt(maxVal, dpMax[i])
	}
	return dpMax[len(dpMax)-1]
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
