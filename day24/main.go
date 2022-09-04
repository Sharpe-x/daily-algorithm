package main

import "container/list"

// 队列 非并发安全
type MyQueue struct {
	list *list.List
}

func Constructor() MyQueue {
	return MyQueue{
		list: list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.list.PushBack(x)
}

func (this *MyQueue) Pop() int {
	e := this.list.Front()
	num := e.Value.(int)
	this.list.Remove(e)
	return num
}

func (this *MyQueue) Peek() int {
	return this.list.Front().Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.list.Len() == 0
}
