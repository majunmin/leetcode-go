package leetcode_0993

import . "github.com/majunmin/leetcode-go/common"

type midRes struct {
	level  int
	parent *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	// 位于 同一层 && parent 不相同
	midResx, midResy := dfs(nil, root, x, 0), dfs(nil, root, y, 0)
	return midResx != nil && midResy != nil &&
		midResx.level == midResy.level && midResx.parent != midResy.parent
}

func dfs(parent *TreeNode, root *TreeNode, val int, level int) *midRes {
	if root.Val == val {
		return &midRes{
			level:  level,
			parent: parent,
		}
	}
	if root.Left != nil {
		if res := dfs(root, root.Left, val, level+1); res != nil {
			return res
		}
	}
	if root.Right != nil {
		if res := dfs(root, root.Right, val, level+1); res != nil {
			return res
		}
	}
	return nil
}
