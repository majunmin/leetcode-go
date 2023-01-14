package quick_sort

func sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left int, right int) {
	//
	if left >= right {
		return
	}

	mid := partition(arr, left, right)
	quickSort(arr, left, mid-1)
	quickSort(arr, mid+1, right)
}

// 分隔
func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	cnt := left
	for i := left; i < right; i++ {
		if arr[i] > pivot {
			continue
		}
		// swap arr[i] arr[cnt]
		arr[i], arr[cnt] = arr[cnt], arr[i]
		cnt++
	}

	//swap arr[cnt] arr[right]
	arr[cnt], arr[right] = arr[right], arr[cnt]
	return cnt
}
