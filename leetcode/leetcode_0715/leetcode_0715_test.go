package leetcode_0715

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	rm := Constructor()
	rm.AddRange(10, 20)
	rm.RemoveRange(14, 16)
	fmt.Println(rm.QueryRange(10, 14))
	fmt.Println(rm.QueryRange(13, 15))
	fmt.Println(rm.QueryRange(16, 17))

}
