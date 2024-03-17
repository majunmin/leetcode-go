package leetcode_0061

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	//
	if head == nil {
		return nil
	}
	// 链表要移动的次数: k % len(list)
	// parse param
	var (
		tail  = head
		total = 1
	)
	for tail.Next != nil {
		total++
		tail = tail.Next
	}
	// rotate right k 次等价于 totate left total - k
	k = total - (k % total)
	// 假设一个头结点
	prev := &ListNode{Next: head}
	// 将 prev.Next 移动到 tail末尾
	for i := 0; i < k; i++ {
		tail.Next = prev.Next
		prev.Next = prev.Next.Next
		prev.Next = nil
	}
	return prev.Next
}
