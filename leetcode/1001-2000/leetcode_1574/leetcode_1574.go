package leetcode_1574

// https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
func findLengthOfShortestSubarray(arr []int) int {

	return solution(arr)
}

// 收缩右端，枚举左端点
func solution(arr []int) int {
	var (
		n = len(arr)
		r = n - 1
	)
	for r > 0 && arr[r-1] <= arr[r] {
		r--
	}
	if r == 0 {
		return 0
	}
	result := r
	for l := 0; l < n; l++ {
		for r < n && arr[r] < arr[l] {
			r++
		}
		// arr[r] >= arr[l]
		result = min(result, r-l-1)
		if l+1 < n && arr[l] > arr[l+1] {
			break
		}
	}
	return result
}

// 收缩左端,枚举右端点
func solution2(arr []int) int {
	var (
		n = len(arr)
		l = 0
	)
	for l < n-1 && arr[l+1] >= arr[l] {
		l++
	}
	if l == n-1 {
		return 0
	}
	result := n - l
	for r := n - 1; r >= 0; r-- {
		for l >= 0 && arr[l] >= arr[r] {
			l--
		}
		result = min(result, r-l-1)
		if r-1 >= 0 && arr[r-1] > arr[r] {
			break
		}
	}
	return result
}
