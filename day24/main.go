package main

import "container/list"

// 队列 非并发安全
type MyQueue_List struct {
	list *list.List
}

func Constructor() MyQueue_List {
	return MyQueue_List{
		list: list.New(),
	}
}

func (this *MyQueue_List) Push(x int) {
	this.list.PushBack(x)
}

func (this *MyQueue_List) Pop() int {
	e := this.list.Front()
	num := e.Value.(int)
	this.list.Remove(e)
	return num
}

func (this *MyQueue_List) Peek() int {
	return this.list.Front().Value.(int)
}

func (this *MyQueue_List) Empty() bool {
	return this.list.Len() == 0
}
