package merge_sort

func sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)>>1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// 合并两个有序数组
func merge(arr []int, left, mid, right int) {
	temp := make([]int, 0, right-left+1)
	i, j := left, mid+1
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}
	for j <= right {
		temp = append(temp, arr[j])
		j++
	}

	// 数组copy
	copy(arr[left:right+1], temp)
}
