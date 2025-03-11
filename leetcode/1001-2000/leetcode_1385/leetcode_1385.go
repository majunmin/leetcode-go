package leetcode_1385

import "slices"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// 1. 枚举所有情况, 时间复杂度 O(mn)
	// 2. 其实可以通过二分查找进行优化。
	//   - 将arr2 排序.
	//   - 判断arr1 中元素是否在arr2中是否符合条件?
	//      - 进需要在arr2中找到第一个 >= arr1中元素的值是否符合条件.

	slices.SortFunc(arr2, func(a, b int) int {
		return a - b
	})
	var (
		result int
	)
	for _, num := range arr1 {
		num1, exist := findTarget(arr2, num-d)
		if !exist {
			result++
			continue
		}
		//num2, exist := findTarget(arr2, num+d)
		//if !exist {
		//	continue
		//}
		//if abs(num-num1) < d && abs(num2-num) < d {
		//	result++
		//}
		if num1 > num+d {
			result++
		}
	}
	return result
}

func findTarget(arr2 []int, target int) (int, bool) {
	n2 := len(arr2)
	idx := lowerBoundary(arr2, target)
	if idx < n2 {
		return arr2[idx], true
	}
	return 0, false
}

func lowerBoundary(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
