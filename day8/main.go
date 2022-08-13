//https://leetcode.cn/problems/reverse-linked-list/

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

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverseListRec(head *ListNode) *ListNode {

	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {

	if cur == nil {
		return pre
	}

	next := cur.Next
	cur.Next = pre

	return reverse(cur, next)
}
