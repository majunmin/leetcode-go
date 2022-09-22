package offer_040

//https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/?envType=study-plan&id=lcof
func getLeastNumbers(arr []int, k int) []int {

	// param check , k > 0 && len(arr) > 0 && len(arr) > k
	length := len(arr)
	if k <= 0 {
		return nil
	}
	if length <= k {
		return arr[:]
	}

	// sort nums
	quickSort(arr, 0, length-1)
	// return result
	return arr[:k]
}

func quickSort(arr []int, left, right int) {

	if left >= right {
		return
	}

	mid := pivot(arr, left, right)

	quickSort(arr, left, mid-1)
	quickSort(arr, mid+1, right)
}

func pivot(arr []int, left int, right int) int {
	pv := arr[right]
	cnt := left
	for i := left; i < right; i++ {
		// swap arr cnt i
		if arr[i] < pv {
			arr[cnt], arr[i] = arr[i], arr[cnt]
			cnt++
		}
	}
	// swap arr right cnt
	arr[right], arr[cnt] = arr[cnt], arr[right]
	return cnt
}
