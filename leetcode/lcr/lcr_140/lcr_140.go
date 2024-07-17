package lcr_140

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlan(head *ListNode, cnt int) *ListNode {
	// 倒数第n个节点
	var (
		slow, fast = head, head
	)
	for i := 0; i < cnt; i++ {
		// 节点个数 不足 cnt个
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
