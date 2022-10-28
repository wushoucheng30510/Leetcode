package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA, currentB := headA, headB
	for currentA != currentB {
		if currentA == nil && currentB == nil {
			return nil
		}
		if currentA == nil && currentB != nil {
			currentA = headB
			currentB = currentB.Next
			continue
		} else if currentA != nil && currentB == nil {
			currentA = currentA.Next
			currentB = headA
			continue
		}
		currentA = currentA.Next
		currentB = currentB.Next
	}
	return currentA

}
