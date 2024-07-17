package leetcode_2974

import "math/rand/v2"

// https://leetcode.cn/problems/minimum-number-game/?envType=daily-question&envId=2024-07-13
func numberGame(nums []int) []int {
	sort(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}
func sort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot-1)
	quickSort(nums, pivot+1, r)
}

func partition(nums []int, l int, r int) int {
	var (
		n        = r - l + 1
		pivotIdx = rand.IntN(n)
		pivot    = nums[r]
		cnt      = l
	)
	nums[r], nums[pivotIdx] = nums[pivotIdx], nums[r]
	for i := l; i < r; i++ {
		if nums[i] > pivot {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}
	nums[cnt], nums[r] = nums[r], nums[cnt]
	return cnt
}
