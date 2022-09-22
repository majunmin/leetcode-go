package offer_040

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	nums := []int{1, 4, 6, 8, 2, 4, 7, 9, 10, 3}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
