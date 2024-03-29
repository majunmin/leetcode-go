package merge_sort

func Sort(arr []int) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	mergeSort(arr, left, right)
}

func mergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)>>1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left int, mid int, right int) {
	tmp := make([]int, right-left+1)
	idx := 0
	lidx, ridx := left, mid+1
	for lidx <= mid && ridx <= right {
		if arr[lidx] < arr[ridx] {
			tmp[idx] = arr[lidx]
			lidx++
		} else {
			tmp[idx] = arr[ridx]
			ridx++
		}
		idx++
	}
	for lidx <= mid {
		tmp[idx] = arr[lidx]
		lidx++
		idx++
	}
	for ridx <= right {
		tmp[idx] = arr[ridx]
		ridx++
		idx++
	}
	copy(arr[left:right+1], tmp)
}
