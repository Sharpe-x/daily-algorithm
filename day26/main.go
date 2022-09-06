package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

type MyStack struct {
	list *list.List
}

func Constructor() MyStack {
	return MyStack{
		list: list.New(),
	}
}

func (this *MyStack) Push(x byte) {
	this.list.PushBack(x)
}

func (this *MyStack) Pop() byte {
	e := this.list.Back()
	if e == nil {
		return 0
	}

	val := e.Value.(byte)
	this.list.Remove(e)

	return val
}

func (this *MyStack) Top() byte {

	e := this.list.Back()
	if e == nil {
		return 0
	}

	return e.Value.(byte)
}

func (this *MyStack) Empty() bool {
	return this.list.Len() == 0
}

func removeDuplicates(s string) string {
	bytes := []byte(s)
	stack := Constructor()
	for i := 0; i < len(bytes); i++ {
		needAdd := true
		for !stack.Empty() && bytes[i] == stack.Top() {
			stack.Pop()
			needAdd = false
		}

		if needAdd {
			stack.Push(bytes[i])
		}
	}

	newStack := Constructor()
	for !stack.Empty() {
		newStack.Push(stack.Pop())
	}

	newBytes := make([]byte, 0)
	for !newStack.Empty() {
		newBytes = append(newBytes, newStack.Pop())
	}

	return string(newBytes)
}

func removeDuplicates2(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
