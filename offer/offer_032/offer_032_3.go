package offer_032

//https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/?envType=study-plan&id=lcof
func levelOrder3(root *TreeNode) [][]int {
	var result [][]int

	// param check
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		l := len(queue)
		curPath := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			curPath = append(curPath, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
		n := len(curPath)

		// reverse  queue
		if n > 0 && level&1 == 1 {
			for i := 0; i < n/2; i++ {
				curPath[i], curPath[n-1-i] = curPath[n-1-i], curPath[i]
			}
		}
		result = append(result, curPath)
		level++

	}
	return result

}
