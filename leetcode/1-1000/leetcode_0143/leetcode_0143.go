package leetcode_0143

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode.cn/problems/reorder-list/?company_slug=pinduoduo
func reorderList(head *ListNode) {
	// param check
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	head.Next = reverse(head.Next)
	reorderList(head.Next)
}

func reverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	dummyNode := new(ListNode)
	for node != nil {
		temp := node.Next
		node.Next = dummyNode.Next
		dummyNode.Next = node
		node = temp
	}
	return dummyNode.Next
}

// 利用 数组(其他语言可能是List)， 将每个节点映射为下标.
func reorderList2(head *ListNode) {
	var (
		nodeList = make([]*ListNode, 0)
	)
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	l, r := 0, len(nodeList)-1
	prev := new(ListNode)
	for l < r {
		prev.Next = nodeList[l]
		l++
		prev = nodeList[l]

		prev.Next = nodeList[r]
		r--
		prev = nodeList[r]
	}
	if l == r {
		prev.Next = nodeList[l]
	}
	prev.Next = nil
}
