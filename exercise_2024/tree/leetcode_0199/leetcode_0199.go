package leetcode_0199

import . "github.com/majunmin/leetcode-go/common"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 层序遍历, 把每一层的最后一个节点加入结果列表
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var result []int
	for len(queue) > 0 {
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
		result = append(result, queue[size-1].Val)
		queue = queue[size:]
	}
	return result
}
