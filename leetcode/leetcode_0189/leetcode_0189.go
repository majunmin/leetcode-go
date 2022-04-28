// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-26 23:19
package leetcode_0189

import "fmt"

// https://leetcode-cn.com/problems/rotate-array/
func rotate(nums []int, k int) {

	if len(nums) <= 1 || k <= 0 {
		return
	}

	if len(nums) < k {
		k %= len(nums)
	}

	rotateSolution3(nums, k)

}

// 空间复杂度 为 O(1)
func rotateSolution3(nums []int, k int) {
	reverseArr(nums, 0, len(nums)-1)
	reverseArr(nums, 0, k-1)
	reverseArr(nums, k, len(nums)-1)
}

func reverseArr(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 使用额外的数组,空间复杂度 O(N)
func rotateSolution1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%len(nums)] = nums[i]
	}
	fmt.Println(newNums)
}
