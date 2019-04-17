package leetcode

//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil {
//		return head
//	}
//	var i int
//	for i = 1; head.Next != nil && i <= k; i++ {
//		head = head.Next
//	}
//	if i < k {
//		return head
//	}
//	head := swap()
//	next := head.Next
//	head.Next = reverseKGroup(next.Next, k)
//	next.Next = head
//	return next
//}
//
//func swap(head *ListNode, k int) *ListNode {
//	var dummy, l, t *ListNode
//	l = head
//	for i := 0; i < k; i++ {
//		t = l.Next
//		l.Next = dummy.Next
//		dummy.Next = l
//		l = t
//	}
//	return dummy
//}
