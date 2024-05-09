package leetcode_0589

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/n-ary-tree-preorder-traversal/description/
func preorder(root *Node) []int {
	//var result []int
	//preorderHelper(root, &result)
	//return result
	return stackSolution(root)
}

func stackSolution(root *Node) []int {
	var stack []*Node
	var result []int

	stack = append(stack, root)
	for len(stack) > 0 {
		// 出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)
		for i := len(node.Children) - 1; i <= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return result

}

func preorderHelper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	for _, cNode := range root.Children {
		preorderHelper(cNode, result)
	}
}
