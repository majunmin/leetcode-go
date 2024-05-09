package leetcode_0045

// https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	//正向查找
	end := 0
	steps := 0
	rightMost := 0
	for i := 0; i < len(nums)-1; i++ {
		rightMost = maxInt(rightMost, nums[i]+i)
		if i == end {
			end = rightMost
			steps++
		}
	}
	return steps
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 反向查找
func solution1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
			}
			steps++
		}
	}
	return steps
}
