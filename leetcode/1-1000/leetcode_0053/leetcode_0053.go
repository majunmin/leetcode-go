package leetcode_0053

type Region struct {
	lSum int // 以 l 为左端点的最大和
	rSum int // 以 r 为右端点的最大和
	iSum int // 当前区间和
	mSum int // 当前区间 最大和
}

// https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	return solution2(nums)
}

// 分治解法 - 线段树
// 类似于 `线段树求解最长公共上升子序列问题` 的pushUp操作，
//   - 如何维护区间信息
//   - 如何合并区间信息呢
func solution2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return get(nums, 0, len(nums)-1).mSum
}

func get(nums []int, l int, r int) Region {
	if l == r {
		return Region{lSum: nums[l], rSum: nums[l], iSum: nums[l], mSum: nums[l]}
	}

	mid := l + (r-l)/2
	lRegion := get(nums, l, mid)
	rRegion := get(nums, mid+1, r)
	return pushUp(lRegion, rRegion)
}

func pushUp(lRegion Region, rRegion Region) Region {
	return Region{
		lSum: maxInt(lRegion.lSum, lRegion.iSum+rRegion.lSum),
		rSum: maxInt(rRegion.rSum, rRegion.iSum+lRegion.rSum),
		iSum: lRegion.iSum + rRegion.iSum,
		mSum: maxInt(maxInt(lRegion.mSum, rRegion.mSum), lRegion.rSum+rRegion.lSum),
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划解法.
// 理解 dp[i] 到已当前位置 截止 的  最大子序和.
func dpSolution(nums []int) int {
	// 计算前面的值  对后面值的增益
	if len(nums) == 0 {
		return 0
	}

	maxVal := nums[0]
	dp := make([]int, len(nums))
	// init state
	dp[0] = nums[0]
	// 递推公式
	// dp[i] = max(nums[i], dp[i-1] + nums[i])
	for i := 1; i < len(nums); i++ {
		dp[i] = maxInt(nums[i], nums[i]+dp[i-1])
		maxVal = maxInt(dp[i], maxVal)
	}
	return maxVal
}
