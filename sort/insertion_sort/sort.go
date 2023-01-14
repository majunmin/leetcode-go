package insertion_sort

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if arr[j] <= val {
				break
			}
			arr[j+1] = arr[j]
		}
		// move elems
		arr[j+1] = val
	}
}
