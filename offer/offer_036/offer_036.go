package offer_036

import . "github.com/majunmin/leetcode-go/offer"

// 二叉树的 中序遍历
func treeToDoublyList(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	dummyNode := &TreeNode{}
	pre := dummyNode
	stack := make([]*TreeNode, 0)

	item := node
	for len(stack) > 0 || item != nil {
		for item != nil {
			stack = append(stack, item)
			item = item.Left
		}

		// pop item
		item = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// build doubleList
		pre.Right = item
		item.Left = pre
		pre = pre.Right

		if item.Right != nil {
			item = item.Right
			continue
		}
		item = nil

	}

	return dummyNode.Right
}

func recurSolution(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	head, _ := tree2DoublyList(node)
	return head
}

func tree2DoublyList(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}

	ll, lr := tree2DoublyList(node.Left)
	rl, rr := tree2DoublyList(node.Right)
	if lr == nil && rl == nil {
		return node, node
	}
	if lr == nil {
		node.Right = rl
		rl.Left = node
		return node, rr

	}
	if rl == nil {
		lr.Right = node
		node.Left = lr
		return ll, node
	}

	lr.Right = node
	node.Left = lr
	node.Right = rl
	rl.Left = node

	return ll, rr

}
