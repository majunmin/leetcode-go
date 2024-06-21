// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-21 20:34
package leetcode_0104

import . "github.com/majunmin/leetcode-go/common"

func maxDepth(root *TreeNode) int {
	// terminate
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		maxLevel int
	)
	// 层序遍历
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		maxLevel++
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		queue = queue[size:]
	}
	return maxLevel
}
