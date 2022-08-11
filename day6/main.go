package main

import "fmt"

// https://leetcode.cn/problems/remove-linked-list-elements/
// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	printLinkList(head)

	newHead := removeElements(head, 3)
	printLinkList(newHead)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	dummyHead := &ListNode{
		Next: head,
		Val:  -1,
	}

	cur := dummyHead
	for cur.Next != nil {
		next := cur.Next
		if next.Val == val {
			cur.Next = next.Next
			next.Next = nil
		} else {
			cur = next
		}
	}

	return dummyHead.Next
}

func printLinkList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("NULL")
}
