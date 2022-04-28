// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-17 17:14
package leetcode_0283

func moveZeroes(nums []int) {
	iterSolution(nums)
}

// solution 双指针
func iterSolution(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 指向第一个为0 的位置
	left := 0
	for i := range nums {
		if nums[i] == 0 {
			continue
		}

		if left != i {
			nums[left] = nums[i]
			nums[i] = 0
		}
		left++
	}
}

//双指针 解法  swap
// 左指针 左边 均为 非0
// 右指针到左指针 处 均为 0
func solutionSwap(nums []int) {
	// 标记需要替换的位置
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			// swap left -right 或者  nums[left] = nums[right],nums[right] = 0
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
