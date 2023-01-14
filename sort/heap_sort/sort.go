package heap_sort

// 堆排序
func sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	length := len(arr)

	// 堆化
	for i := length/2 - 1; i >= 0; i-- {
		heapify(arr, length, i)
	}

	//
	for i := length - 1; i >= 1; i-- {
		//	 swap  arr[i], arr[0]
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}

}

func heapify(arr []int, length int, idx int) {
	left, right := 2*idx+1, 2*idx+2
	largestIdx := idx
	if left < length && arr[left] > arr[largestIdx] {
		largestIdx = left
	}
	if right < length && arr[right] > arr[largestIdx] {
		largestIdx = right
	}
	// swap  arr[largestIdx] , arr[idx]
	if largestIdx != idx {
		arr[idx], arr[largestIdx] = arr[largestIdx], arr[idx]
		heapify(arr, length, largestIdx)
	}
}
