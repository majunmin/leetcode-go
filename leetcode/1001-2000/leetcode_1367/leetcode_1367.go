package leetcode_1367

import (
	. "github.com/majunmin/leetcode-go/common"
)

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool {
	// 链表完全匹配
	if head == nil {
		return true
	}
	// 二叉树匹配到空节点
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
}
