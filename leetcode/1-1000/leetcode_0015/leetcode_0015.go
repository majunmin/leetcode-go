// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-18 21:35
package leetcode_0015

import "sort"

//https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i, num := range nums {
		if num > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right]+num == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				// 去重
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right]+num < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
