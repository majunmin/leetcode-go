package leetcode_0114

import (
	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/?envType=study-plan-v2&envId=top-100-liked
func flatten(root *TreeNode) {
	// param check
	var preorder []*TreeNode
	node := root
	var stack []*TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			preorder = append(preorder, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	// 串成单链表
	preNode := &TreeNode{}
	for _, node = range preorder {
		preNode.Right = node
		preNode.Left = nil
		preNode = node
	}
}

// 递归的方式, inplace 解决
func flatten2(root *TreeNode) {
	// 把 二叉树的 左节点打平， 然后插入到右节点下
	// terminate
	if root == nil {
		return
	}
	flatten2(root.Left)
	flatten2(root.Right)
	// 拼接链表
	tmp := root.Right
	root.Right, root.Left = root.Left, nil
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = tmp
}

// 二叉树的后续遍历(迭代方式)-- 自底向上的解决方案
func flatten3(root *TreeNode) {
	// 把 二叉树的 左节点打平， 然后插入到右节点下
	// terminate
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)

	// 用于构建单链表
	var prev *TreeNode
	// 用于二叉树后续遍历
	var node, pre *TreeNode
	node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}
		// pop stack
		node = stack[len(stack)-1]
		if node.Left == nil || node.Left == pre {
			stack = stack[:len(stack)-1]
			// process
			node.Left, node.Right = nil, prev
			prev = node
			//res.aa(node)
			pre = node
			node = nil
			continue
		}
		node = node.Left
	}
}
