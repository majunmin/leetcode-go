// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-21 20:34
package common

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node N叉树节点
type Node struct {
	Val      int
	Children []*Node
}
