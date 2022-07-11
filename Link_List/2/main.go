package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	dummy := &ListNode{}
	tmp := dummy

	// O(max(len(l1),len(l2))+1)
	for l1 != nil || l2 != nil || sum > 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode{Val: sum % 10, Next: nil}
		tmp = tmp.Next
		sum /= 10
	}
	return dummy.Next
}
