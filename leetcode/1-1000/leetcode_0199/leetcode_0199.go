package leetcode_0199

import (
	. "github.com/majunmin/leetcode-go/common"
)

func rightSideView(root *TreeNode) []int {
	// 利用 二叉树的 层次遍历, 从右侧开始遍历
	// param check
	if root == nil {
		return nil
	}
	//
	var result []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		// add to result
		result = append(result, queue[len(queue)-1].Val)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		queue = queue[size:]
	}
	return result
}
