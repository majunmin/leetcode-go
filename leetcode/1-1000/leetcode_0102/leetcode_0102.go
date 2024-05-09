package leetcode_0102

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/binary-tree-level-order-traversal/#/description
func levelOrder(root *TreeNode) [][]int {
	// param check
	if root == nil {
		return nil
	}
	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0, len(queue))
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			node := queue[i]
			// process cur node
			level = append(level, node.Val)

			// add queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[curLen:]
		result = append(result, level)
	}
	return result
}
