package leetcode_1261

import . "github.com/majunmin/leetcode-go/common"

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	transVerseTree(root, 0)
	return FindElements{root: root}
}

func transVerseTree(node *TreeNode, val int) {
	if node == nil {
		return
	}
	node.Val = val
	transVerseTree(node.Left, val*2+1)
	transVerseTree(node.Right, val*2+2)
}

func (this *FindElements) Find(target int) bool {

	node := find(target, this.root)
	return node != nil
}

// find target 所在的节点
func find(target int, node *TreeNode) *TreeNode {
	// terminate
	if node == nil {
		return nil
	}
	if target == 0 && node != nil {
		return node
	}
	//处理当前层
	parent := find((target-1)>>1, node)
	// 左子节点 是奇数， 右子节点是偶数
	if target&1 == 1 {
		return parent.Left
	} else {
		return parent.Right
	}
}
