package offer_016

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(0.00001, 2147483647))
}
