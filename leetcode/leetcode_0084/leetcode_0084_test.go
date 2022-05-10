package leetcode_0084

import (
	"fmt"
	"testing"
)

func TestMaxRectangle(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{5, 4, 3, 2}))
	fmt.Println(largestRectangleArea([]int{2, 3, 4, 5}))
}
