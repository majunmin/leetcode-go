package leetcode_1122

// https://leetcode.cn/problems/relative-sort-array/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	// pram check
	if len(arr1) == 0 {
		return nil
	}
	var idx int
	// sort by arr2
	for _, target := range arr2 {
		for i := idx; i < len(arr1); i++ {
			if arr1[i] == target {
				arr1[i], arr1[idx] = arr1[idx], arr1[i]
				idx++
			}
		}
	}

	// sort  remain  idx ~ (len(arr1)-1)
	sort(arr1, idx, len(arr1))
	return arr1

}

func sort(arr []int, start int, end int) {
	// 冒泡排序
	//for i := start; i < end; i++ {
	//	for j := i + 1; j < end; j++ {
	//		if arr[j] < arr[i] {
	//			arr[i], arr[j] = arr[j], arr[i]
	//		}
	//	}
	//}

	// 插入排序
	for i := start; i < end-1; i++ {
		minVal := arr[i]
		right := i
		for j := i; j < end; j++ {
			if arr[j] < minVal {
				minVal = arr[j]
				right = j
			}
		}
		// swap left & right
		arr[i], arr[right] = arr[right], arr[i]
	}
}
