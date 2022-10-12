package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	inputList := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{5, nil}}}}}
	outputList := deleteDuplicates(inputList)
	for outputList != nil {
		fmt.Println(outputList.Val)
		outputList = outputList.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preList, nextList := head, head.Next
	for nextList != nil {
		if preList.Val == nextList.Val {
			nextList = nextList.Next
			continue
		} else {
			preList.Next = nextList
			preList = preList.Next
			nextList = nextList.Next
		}
	}
	if nextList == nil && preList.Next != nil {
		preList.Next = nil
	}
	return head
}
