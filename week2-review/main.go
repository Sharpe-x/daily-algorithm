package main

import "fmt"

func main() {
	head := getList([]int{1, 2, 3, 4, 5})
	printList(head)
	newHead := removeElements(head, 3)
	printList(newHead)

	head2 := getList([]int{1, 2, 3, 4, 5})
	rhead2 := reverseList(head2)
	printList(rhead2)

	head3 := getList([]int{1, 2, 3, 4, 5})
	rhead3 := reverseList(head3)
	printList(rhead3)
}

func removeElements(head *ListNode, val int) *ListNode {

	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	cur := dummyHead
	for cur != nil {

		next := cur.Next
		for next != nil && next.Val == val {
			cur.Next = next.Next
			next.Next = nil
			next = cur.Next
		}

		cur = next
	}

	return dummyHead.Next
}

func getList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		cur = cur.Next
	}

	return head
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {

		fmt.Printf("%d - > ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("NULL")
}

//
// 设计链表的实现。
// 您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。
// val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
// 在链表类中实现这些功能：
// get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。
// 插入后，新节点将成为链表的第一个节点。
// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
// addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
// deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	dummyHead *ListNode
	size      int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &ListNode{
			Val:  -1,
			Next: nil,
		},
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 {
		return -1
	}

	cur := this.dummyHead
	for index >= 0 {
		cur = cur.Next
		index--
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {

	cur := this.dummyHead
	next := cur.Next
	cur.Next = &ListNode{
		Val:  val,
		Next: next,
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.dummyHead
	n := this.size
	for n > 0 {
		cur = cur.Next
		n--
	}
	cur.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	if index < 0 {
		this.AddAtHead(val)
		return
	}

	cur := this.dummyHead
	for index > 0 {
		cur = cur.Next
		index--
	}
	next := cur.Next
	cur.Next = &ListNode{
		Val:  val,
		Next: next,
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}

	cur := this.dummyHead
	for index > 0 {
		cur = cur.Next
		index--
	}
	del := cur.Next
	cur.Next = del.Next
	del.Next = nil
	this.size--
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

func swapPairs(head *ListNode) *ListNode {

	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	pre := dummyHead
	for pre.Next != nil && pre.Next.Next != nil {
		node1, node2 := pre.Next, pre.Next.Next
		next := node2.Next

		pre.Next = node2
		node2.Next = node1
		node1.Next = next

		pre = node1
	}

	return dummyHead.Next
}

func swapPairsRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairsRec(next.Next)
	next.Next = head

	return next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: head,
	}

	slow, fast := dummyHead, dummyHead
	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	del := slow.Next
	next := del.Next

	slow.Next = next
	del.Next = nil

	return dummyHead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA int
	curA := headA
	for curA != nil {
		lenA++
		curA = curA.Next
	}

	var lenB int
	curB := headB
	for curB != nil {
		lenB++
		curB = curB.Next
	}

	curA, curB = headA, headB
	if lenA > lenB {
		for lenA-lenB > 0 {
			curA = curA.Next
			lenA--
		}
	} else {
		for lenB-lenA > 0 {
			curB = curB.Next
			lenB--
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
