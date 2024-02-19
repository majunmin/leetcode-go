package leetcode_0845

// https://leetcode.cn/problems/longest-mountain-in-array/
func longestMountain(arr []int) int {
	// param check
	if len(arr) < 3 {
		return 0
	}
	size := len(arr)
	left, right := make([]int, size), make([]int, size)
	// left[i] 左边有多少连续的左山脉
	// right[i] 右边有多少连续的右山脉
	// init state
	left[0], right[size-1] = 0, 0
	for i := 1; i < size-1; i++ {
		if arr[i] > arr[i-1] {
			left[i] = left[i-1] + 1
		}
		if arr[size-1-i] > arr[size-i] {
			right[size-1-i] = right[size-i] + 1
		}
	}
	//
	var result int
	for i := 1; i < size-1; i++ {
		result = max(result, left[i]+right[i])
	}

	return result
}
