package leetcode_0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/?envType=study-plan-v2&envId=top-100-liked
func copyRandomList(head *Node) *Node {
	p := head
	for p != nil {
		newNode := &Node{Val: p.Val}
		newNode.Next = p.Next
		p.Next = newNode
		p = p.Next.Next
	}

	// 构建 random 链表
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// 将 复制链表拆出来
	dummyNode := &Node{}
	prev := dummyNode
	p = head
	for p != nil {
		prev.Next = p.Next
		p.Next = p.Next.Next
		prev = prev.Next
		p = p.Next
	}
	return dummyNode.Next
}

var (
	nodeCache = make(map[*Node]*Node)
)

// 递归解法
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	// process

	if _, exist := nodeCache[head]; !exist {
		newNode := &Node{
			Val: head.Val,
		}
		nodeCache[head] = newNode
		newNode.Next = copyRandomList2(head.Next)
		newNode.Random = copyRandomList2(head.Random)
	}
	return nodeCache[head]
}

func copyRandomList1(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	p := head
	for p != nil {
		nodeMap[p] = &Node{
			Val: p.Val,
		}
		p = p.Next
	}
	p = head
	for p != nil {
		cpyNode := nodeMap[p]
		cpyNode.Next = nodeMap[head.Next]
		cpyNode.Random = nodeMap[head.Random]
		p = p.Next
	}
	return nodeMap[head]
}
