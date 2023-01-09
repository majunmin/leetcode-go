package offer_059

import (
	"fmt"
	"testing"
)

func Test_MaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}
