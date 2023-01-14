package bubble_sort

// stable sort
// 冒泡排序
func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				//swap arr[i] arr [j]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
