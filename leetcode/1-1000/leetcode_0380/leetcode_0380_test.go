package leetcode_0380

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	randSet := Constructor()
	randSet.Insert(1)
	randSet.Remove(2)
	randSet.Insert(2)
	fmt.Println(randSet.GetRandom())
}
