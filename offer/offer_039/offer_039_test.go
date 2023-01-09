package offer_039

import (
	"fmt"
	"testing"
)

func Test_MahorNum(t *testing.T) {
	fmt.Println(majorityElement([]int{6, 5, 5}))
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
