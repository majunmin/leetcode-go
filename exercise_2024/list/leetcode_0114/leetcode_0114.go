package leetcode_0114

import . "github.com/majunmin/leetcode-go/common"

var pre *TreeNode

func flatten(root *TreeNode) {
	// 把 二叉树的 左节点打平， 然后插入到右节点下
	// terminate
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 拼接链表
	tmp := root.Right
	root.Right, root.Left = root.Left, nil
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = tmp
}
