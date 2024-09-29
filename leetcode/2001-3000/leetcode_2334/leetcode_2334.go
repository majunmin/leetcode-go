package leetcode_2334

import "fmt"

// https://leetcode.cn/problems/subarray-with-elements-greater-than-varying-threshold/
func validSubarraySize(nums []int, threshold int) int {
	// 题目求解   nums中长为k的子数组, min(nums) * k <= threshold.
	// 可以转化为 leetcode85(求解矩形面积).
	//
	// 单调栈, 添加哨兵,简化计算.
	nums = append([]int{0}, nums...)
	nums = append(nums, 0)
	//
	var (
		n     = len(nums)
		stack = make([]int, 0)
	)
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
			// pop stack
			h := nums[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			if h*w > threshold {
				fmt.Println(h, "-", w)
				return w
			}
		}
		stack = append(stack, i)
	}
	return -1
}
