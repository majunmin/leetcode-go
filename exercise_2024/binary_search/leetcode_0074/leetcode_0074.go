package leetcode_0074

// https://leetcode.cn/problems/search-a-2d-matrix/?envType=study-plan-v2&envId=top-100-liked
func searchMatrix(matrix [][]int, target int) bool {
	// 两次二分法
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	//
	col := findCol(matrix, target)
	if col == 0 {
		return false
	}
	if matrix[col][0] == target {
		return true
	}
	idx := findIdx(matrix[col-1], target)
	return matrix[col-1][idx] == target
}

func findIdx(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func findCol(matrix [][]int, target int) int {
	left, right := 0, len(matrix)-1
	for left < right {
		mid := left + (right-left)>>1
		if matrix[mid][0] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
