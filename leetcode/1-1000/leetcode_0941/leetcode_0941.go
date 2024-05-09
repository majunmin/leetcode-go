package leetcode_0941

// https://leetcode.cn/problems/valid-mountain-array/
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	idx := 0
	for idx < len(arr)-1 {
		if arr[idx] >= arr[idx+1] {
			break
		}
		idx++
	}
	if idx == 0 {
		return false
	}
	// 比较右侧
	for idx < len(arr)-1 {
		if arr[idx] <= arr[idx+1] {
			break
		}
		idx++
	}

	return idx == len(arr)-1
}

func solution1(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	for i := 1; i < len(arr)-1; i++ {
		left, right := i, i
		for left > 0 {
			if arr[left] <= arr[left-1] {
				break
			}
			left--
		}
		for right < len(arr)-1 {
			if arr[right] <= arr[right+1] {
				break
			}
			right++
		}
		if left == 0 && right == len(arr)-1 {
			return true
		}
	}
	return false
}
