package leetcode_0589

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	// param check
	if root == nil {
		return nil
	}
	stack := make([]*Node, 0)
	var result []int
	stack = append(stack, root)
	for len(stack) > 0 {
		// pop stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, item.Val)
		for i := len(root.Children); i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}
	return result
}
