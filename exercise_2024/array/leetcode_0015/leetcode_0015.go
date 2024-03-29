package leetcode_0015

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//
		left, right := i+1, len(nums)-1
		for left < right {
			for left < right && nums[left]+nums[right]+nums[i] > 0 {
				right--
			}
			for left < right && nums[left]+nums[right]+nums[i] < 0 {
				left++
			}
			if left < right && nums[left]+nums[right]+nums[i] == 0 {
				result = append(result, []int{nums[left], nums[right], nums[i]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}
