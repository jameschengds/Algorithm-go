package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var del *ListNode
	node := head
	if head.Next == nil {
		if n == 0 {
			return head
		} else if n == 1 {
			return nil
		}
	}
	for i := 0; node.Next != nil; i++ {
		if i == n-1 {
			del = head
		}
		node = node.Next
		if i >= n && del != nil {
			del = del.Next
		}
	}
	if del == nil {
		head = head.Next
		return head
	}
	if del.Next.Next != nil {
		del.Next = del.Next.Next
	} else {
		del.Next = nil
	}

	return head
}
