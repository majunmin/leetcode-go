package insertion_sort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{1, 2, 5, 4, 3}
	sort(arr)
	fmt.Println(arr)
}
