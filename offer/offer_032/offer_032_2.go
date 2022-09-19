package offer_032

func levelOrder2(root *TreeNode) [][]int {
	var result [][]int
	// param check
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		var curLevel []int
		for i := 0; i < l; i++ {
			node := queue[i]

			// process
			curLevel = append(curLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, curLevel)
		queue = queue[l:]
	}
	return result

}
