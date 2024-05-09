package leetcode_0443

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {

	fmt.Println(minMutation("AACCGGTT", "AAACGGTA",
		[]string{"AACCGGTA", "AAACGGTT", "AAACGGTT", "AAACGGTA"}))
}
