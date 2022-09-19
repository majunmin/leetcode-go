package offer_011

import (
	"fmt"
	"testing"
)

func Test_MinArray(t *testing.T) {
	fmt.Println(minArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
	fmt.Println(minArray([]int{2, 2, 2, 0}))
	fmt.Println(minArray([]int{10, 1, 10, 10, 10}))
}
