package offer_053

import (
	"fmt"
	"testing"
)

func Test_MissingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{0, 2}))
	fmt.Println(missingNumber([]int{1, 2}))
}
