// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 21:27
package leetcode_0101

import . "github.com/majunmin/leetcode-go/common"

func isSymmetric(root *TreeNode) bool {
	// 方法二 迭代形式, 二叉树的层序遍历.
	if root == nil {
		return false
	}
	var (
		queue1 = make([]*TreeNode, 0)
		queue2 = make([]*TreeNode, 0)
	)
	queue1 = append(queue1, root)
	queue2 = append(queue2, root)
	for len(queue1) > 0 && len(queue2) > 0 {
		if len(queue1) != len(queue2) {
			return false
		}
		size := len(queue1)
		for i := 0; i < size; i++ {
			node1 := queue1[i]
			node2 := queue2[i]
			if node1 == nil && node2 == nil {
				continue
			}
			if node1 == nil || node2 == nil || node1.Val != node2.Val {
				return false
			}
			queue1 = append(queue1, node1.Left)
			queue1 = append(queue1, node1.Right)
			queue2 = append(queue2, node2.Right)
			queue2 = append(queue2, node2.Left)
		}
		queue1 = queue1[size:]
		queue2 = queue2[size:]
	}
	return true
}

// 1. 递归形式解法
func recursiveSolution(root *TreeNode) bool {
	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isSymmetricHelper(node1.Left, node2.Right) && isSymmetricHelper(node1.Right, node2.Left)
}
