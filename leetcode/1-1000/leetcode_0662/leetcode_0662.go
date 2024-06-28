package leetcode_0662

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type wrapNode struct {
	index int
	node  *TreeNode
}

func newWrapNode(index int, node *TreeNode) wrapNode {
	return wrapNode{
		index: index,
		node:  node,
	}
}

// https://leetcode.cn/problems/maximum-width-of-binary-tree/
func widthOfBinaryTree(root *TreeNode) int {
	var (
		queue  = make([]wrapNode, 0)
		result int
	)
	queue = append(queue, newWrapNode(1, root))
	for len(queue) > 0 {
		size := len(queue)
		// update result
		if size > 1 {
			left, right := queue[0].index, queue[len(queue)-1].index
			result = max(result, right-left+1)
		}
		for i := 0; i < size; i++ {
			item := queue[i]
			idx, node := item.index, item.node
			if node.Left != nil {
				queue = append(queue, newWrapNode(idx*2, node.Left))
			}
			if node.Right != nil {
				queue = append(queue, newWrapNode(idx*2+1, node.Right))
			}
		}
		queue = queue[size:]
	}
	return result
}
