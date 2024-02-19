package leetcode_0852

func peakIndexInMountainArray(arr []int) int {
	// 二分查找
	// param check
	if len(arr) < 3 {
		panic("invalid param")
	}
	left, right := 1, len(arr)-2
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] > arr[mid+1] {
			right = mid
		} else if arr[mid] > arr[mid-1] {
			left = mid + 1
		}
	}
	return left
}
