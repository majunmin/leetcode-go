package lcr_077

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/7WHec2/
func sortList(head *ListNode) *ListNode {
	// 插入排序 or 归并排序
	if head == nil || head.Next == nil {
		return head
	}
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// find mid
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	l1 := mergeSort(head)
	l2 := mergeSort(mid)
	return merge(l1, l2)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}

func insertSort(head *ListNode) *ListNode {
	var (
		dummyNode = new(ListNode)
		cur       = head
	)
	for cur != nil {
		pre := dummyNode
		tmp := cur.Next
		for pre.Next != nil {
			if pre.Next.Val < cur.Val {
				break
			}
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		cur = tmp
	}
	return dummyNode.Next
}
