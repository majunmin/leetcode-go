package leetcode_0103

import (
	"slices"

	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
	// param check
	if root == nil {
		return nil
	}
	var result [][]int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var levelIdx int
	for len(queue) > 0 {
		size := len(queue)
		var curLevel []int
		for i := 0; i < size; i++ {
			item := queue[i]
			curLevel = append(curLevel, item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		queue = queue[size:]
		if levelIdx&1 == 1 {
			slices.Reverse(curLevel)
		}
		levelIdx++
		slices.Reverse(curLevel)
		result = append(result, curLevel)
	}
	return result
}
