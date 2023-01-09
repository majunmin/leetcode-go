package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	n := len(nums)
	for i := 0; i < n-2; i++ {
		target := -nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			if nums[left]+nums[right] < target {
				left++
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}

	}
	return result
}
