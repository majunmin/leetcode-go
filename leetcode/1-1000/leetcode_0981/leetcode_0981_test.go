package leetcode_0981

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	fmt.Println(timeMap.Get("foo", 1))
}
