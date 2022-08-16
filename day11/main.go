package main

// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	lenA := 0
	for curA != nil {
		curA = curA.Next
		lenA++
	}

	curB := headB
	lenB := 0
	for curB != nil {
		curB = curB.Next
		lenB++
	}

	step := 0
	curA, curB = headA, headB
	if lenA > lenB {
		step = lenA - lenB
		for step > 0 {
			curA = curA.Next
			step--
		}
	} else {
		step = lenB - lenA
		for step > 0 {
			curB = curB.Next
			step--
		}
	}

	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}

	return nil
}
