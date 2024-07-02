package lcr_069

// https://leetcode.cn/problems/B1IidL/
func peakIndexInMountainArray(arr []int) int {
	// O（N）
	var (
		result int
	)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			result++
		}
	}
	return result
}

// binary search
func peakIndexInMountainArray2(arr []int) int {

	l, r := -1, len(arr)
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid] > arr[mid+1] {
			r = mid
		} else {
			l = mid
		}
	}
	return l + 1
}
