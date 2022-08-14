package main

// https://leetcode.cn/problems/swap-nodes-in-pairs/submissions/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		p, q := cur.Next, cur.Next.Next
		next := q.Next

		q.Next = p
		p.Next = next
		cur.Next = q
		cur = p
	}

	return dummyHead.Next
}
