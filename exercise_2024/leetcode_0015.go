package exercise_2024

import (
	"sort"
)

// https://leetcode.cn/problems/3sum/?envType=study-plan-v2&envId=top-100-liked
func threeSum(nums []int) [][]int {
	// param check
	if len(nums) < 3 {
		return nil
	}
	// process
	sort.Ints(nums)
	n := len(nums)
	var result [][]int
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 || (i > 0 && nums[i] == nums[i-1]) {
			break
		}
		result = append(result, findTarget(nums, i+1, n-1, nums[i])...)
	}
	return result
}

func findTarget(nums []int, left, right, num int) [][]int {
	var result [][]int
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
		} else if nums[left]+nums[right]+num > 0 {
			right--
		} else {
			left++
		}
	}
	return result
}
