package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here

	var newHead *ListNode
	for pHead != nil {
		next := pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = next
	}

	return newHead
}
