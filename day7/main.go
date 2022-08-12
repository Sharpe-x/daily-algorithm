// https://leetcode.cn/problems/design-linked-list/
package main

func main() {

}

type MyLinkedList struct {
	dummy *Node
	size  int
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	Node := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	return MyLinkedList{dummy: Node, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size-1 {
		return -1
	}

	cur := this.dummy
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {

	curHead := this.dummy.Next
	newNode := &Node{
		Val:  val,
		Next: curHead,
	}

	this.dummy.Next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {

	cur := this.dummy
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &Node{
		Val:  val,
		Next: nil,
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	} else if index == this.size {
		this.AddAtTail(val)
		return
	} else if index > this.size {
		return
	}

	node := Node{
		Val:  val,
		Next: nil,
	}

	cur := this.dummy
	for index > 0 {
		index--
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = &node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}

	cur := this.dummy
	for index > 0 {
		index--
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
