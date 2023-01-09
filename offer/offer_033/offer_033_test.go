package offer_033

import (
	"fmt"
	"testing"
)

func Test_VerifyPostOrder(t *testing.T) {
	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
	fmt.Println(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}))
}
