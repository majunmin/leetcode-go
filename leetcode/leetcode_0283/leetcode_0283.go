// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-17 17:14
package leetcode_0283

//双指针 解法
// 左指针 左边 均为 非0
// 右指针到左指针 处 均为 0
func moveZeroes(nums []int) {
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
