package offer_007

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {

	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree)

}
