package _023_20_22

import (
	"math"
)

func minimumSum(nums []int) int {
	// param check
	if len(nums) < 3 {
		return -1
	}
	result := math.MaxInt
	for i := 1; i < len(nums)-1; i++ {
		result = minInt(result, minSum(nums, i))
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}

// 计算以 i 为 山顶的 最小和
func minSum(nums []int, mid int) int {
	left, right := mid-1, mid+1
	minLeft, minRight := nums[mid], nums[mid]
	for left >= 0 {
		if nums[left] < minLeft {
			minLeft = nums[left]
		}
		left--
	}
	for right < len(nums) {
		if nums[right] < minRight {
			minRight = nums[right]
		}
		right++
	}
	if minLeft == nums[mid] || minRight == nums[mid] {
		return math.MaxInt
	}
	return minLeft + minRight + nums[mid]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minInt3(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func validMountain(nums []int, i, j, k int) bool {
	return nums[i] < nums[j] && nums[j] > nums[k]
}
