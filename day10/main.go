package main

func main() {

}

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
	dummHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	fast, slow := dummHead, dummHead
	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	del := slow.Next
	slow.Next = del.Next
	del.Next = nil

	return dummHead.Next
}
