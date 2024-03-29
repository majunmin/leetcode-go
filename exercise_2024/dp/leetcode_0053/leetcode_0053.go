package leetcode_0053

// https: //leetcode.cn/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	return partitionSolution(nums)
}

type midResult struct {
	lMax   int
	rMax   int
	curSum int
	curMax int
}

func partitionSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// process
	res := partition(nums, 0, len(nums)-1)
	return res.curMax

}

func partition(nums []int, left, right int) midResult {
	//terminate
	if left == right {
		return midResult{
			lMax:   nums[left],
			rMax:   nums[left],
			curSum: nums[left],
			curMax: nums[left],
		}
	}
	// process
	mid := left + (right-left)>>1
	lRes, rRes := partition(nums, left, mid), partition(nums, mid+1, right)
	return midResult{
		lMax:   max(lRes.lMax, lRes.curSum+rRes.lMax),
		rMax:   max(rRes.rMax, rRes.curSum+lRes.rMax),
		curSum: lRes.curSum + rRes.curSum,
		curMax: max(lRes.curMax, rRes.curMax, lRes.rMax+rRes.lMax),
	}
}

func dpSolution(nums []int) int {
	// param check
	if len(nums) < 1 {
		return 0
	}
	//dp[i] : nums数组以 第 i 个位置为止 对后面数组和的 最大增益
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = max(dp[i-1]+nums[i-1], 0)
	}
	var result int
	for i := 1; i < len(dp); i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
