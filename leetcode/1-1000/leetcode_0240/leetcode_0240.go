package leetcode_0240

// https://leetcode.cn/problems/search-a-2d-matrix-ii/?envType=study-plan-v2&envId=top-100-liked
func searchMatrix(matrix [][]int, target int) bool {
	return solution2(matrix, target)
}

// 比较巧妙的方法
func solution2(matrix [][]int, target int) bool {
	rowSize, colSize := len(matrix), len(matrix[0])
	m, n := 0, colSize-1
	for m < rowSize && n >= 0 {
		if matrix[m][n] > target {
			n--
		} else if matrix[m][n] < target {
			m++
		} else {
			return true
		}
	}
	return false
}

// 逐行进行二分查找
func solution1(matrix [][]int, target int) bool {
	// param check
	rowSize, colSize := len(matrix), len(matrix[0])
	for i := 0; i < rowSize; i++ {
		if matrix[i][0] > target {
			return false
		}
		if matrix[i][colSize-1] < target {
			continue
		}
		// binarySearch
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

// 找到右边第一个 <= target 的元素
func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)>>1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return nums[left] == target
}
