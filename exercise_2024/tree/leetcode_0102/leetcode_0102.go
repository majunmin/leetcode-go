package leetcode_0102

import . "github.com/majunmin/leetcode-go/common"

func levelOrder(root *TreeNode) [][]int {
	// param check
	if root == nil {
		return nil
	}
	//
	var (
		result [][]int
		queue  = make([]*TreeNode, 0)
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			item := queue[i]
			level = append(level, item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		result = append(result, level)
		queue = queue[size:]
	}
	return result
}
