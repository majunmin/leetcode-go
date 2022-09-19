package offer_006

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	// param check
	if head == nil {
		return []int{}
	}

	return append(reversePrint(head.Next), head.Val)
}
