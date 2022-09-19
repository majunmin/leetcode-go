package offer_053

import (
	"fmt"
	"testing"
)

func Test_Search(t *testing.T) {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 7))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 10))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 11))
}
