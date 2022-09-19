package offer_022

import . "github.com/majunmin/leetcode-go/offer"

//https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/?envType=study-plan&id=lcof
func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	tail := dummyNode
	for i := 0; i <= k; i++ {
		if tail == nil {
			return nil
		}
		tail = tail.Next
	}
	pre := dummyNode
	for tail != nil {
		tail = tail.Next
		pre = pre.Next
	}
	// delete  kth  Node from end
	//pre.Next = pre.Next.Next
	return pre.Next
}
