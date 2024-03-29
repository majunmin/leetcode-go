package leetcode_0041

func firstMissingPositive(nums []int) int {
	return solution2(nums)
}

// 原地hash 交换位置.  需要掌握o
func solution2(nums []int) int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		// nums[i] > 0 && nums[i] <= n  合法性验证
		// nums[i-1] == i 说明该位置是正确的,没必要进行交换
		for nums[i] > 0 && nums[i] <= n && nums[i-1] != i {
			// swap nums[i]  nums[nums[i-1]]
			nums[i], nums[nums[i]-1] = nums[nums[i-1]], nums[i]
		}
	}
	//
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

// 原地hash, 负数标记,(想不到啊想不到). 理解了就好
func solution1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	// hash 表
	for i := 0; i < n; i++ {
		//  避免 后面的数字 被该流程转负数的流程影响
		num := abs(nums[i])
		if num-1 <= 0 || abs(nums[i-1]) > n {
			continue
		}
		nums[num-1] = -(abs(nums[num-1]))
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
