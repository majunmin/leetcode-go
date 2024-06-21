package leetcode_0117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/?envType=study-plan-v2&envId=top-interview-150
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 二叉树的层序遍历
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		prev := queue[0]
		for i := 0; i < size; i++ {
			cur := queue[i]
			if i > 0 {
				prev.Next = cur
				prev = cur
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[size:]
	}
	return root
}
