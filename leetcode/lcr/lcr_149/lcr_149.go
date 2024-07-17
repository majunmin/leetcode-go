package lcr_149

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func decorateRecord(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// BFS
	var (
		queue  = make([]*TreeNode, 0)
		result = make([]int, 0)
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			//add to Result
			result = append(result, item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		queue = queue[size:]
	}
	return result
}
