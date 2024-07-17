package leetcode_3011

import "slices"

// https://leetcode.cn/problems/find-if-array-can-be-sorted/
// 这一题目的类型是 或者编码技巧是 `分组循环`
// 1. 将数组按照 bitCount 子数组分组, 在每个组内进行排序
// 在遍历数组 判断数组是否整体有序 -- 相当于给数组拍了个序
// 2. 与方法一类似,在每个分组内不必排序,只需记录[最小值,最大值]即可.
func canSortArray(nums []int) bool {
	// solution2 分组排序 的优化， 仅记录最大值最小值
	var (
		i    int
		pMax = -1
	)
	for i < len(nums) {
		var (
			ones           = bitCount(nums[i])
			minVal, maxVal = nums[i], nums[i]
		)
		for ; i < len(nums) && bitCount(nums[i]) == ones; i++ {
			minVal = min(minVal, nums[i])
			maxVal = max(maxVal, nums[i])
		}
		// [start, i)
		if minVal < pMax {
			return false
		}
		pMax = maxVal
	}
	return true
}

// 分组排序
func solution1(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		start := i
		ones := bitCount(nums[i])
		for ; i < len(nums) && bitCount(nums[i]) == ones; i++ {
			slices.Sort(nums[start:i])
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func bitCount(n int) int {
	var res int
	for n > 0 {
		res++
		n = n & (n - 1) // 将 2进制 低位置为0
	}
	return res
}

// 2024-07-13 自己想的做的 暴力解法, 类似于 冒泡排序法进行排序,
// - 每次迭代将最大值移到最后面, 不能成功就返回false
func solution(nums []int) bool {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] <= nums[j+1] {
				continue
			}
			if bitCount(nums[j]) != bitCount(nums[j+1]) {
				return false
			}
			// swap
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
	return true
}
