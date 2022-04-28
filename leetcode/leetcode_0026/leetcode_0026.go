// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-26 23:10
package leetcode_0026

//要求 原地删除 ,空间复杂度为 O(1)
//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	//双指针法
	// 判断重复项 nums[i] == nums[i-1]
	i := 1
	for j := 1; j < len(nums); j++ {
		// 重复项
		if nums[j] == nums[j-1] {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	return i
}
