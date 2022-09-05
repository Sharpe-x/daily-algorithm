package main

import "container/list"

type MyQueue struct {
	stack []int
	back  []int
}

func ConstructorQueue() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.back) == 0 {
		for len(this.stack) != 0 {
			val := this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
			this.back = append(this.back, val)
		}
	}

	if len(this.back) == 0 {
		return 0
	}

	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == 0 {
		return 0
	}
	this.back = append(this.back, val)
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}

type MyStack struct {
	list *list.List
}

func Constructor() MyStack {
	return MyStack{
		list: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.list.PushBack(x)
}

func (this *MyStack) Pop() int {
	e := this.list.Back()
	if e == nil {
		return 0
	}

	val := e.Value.(int)
	this.list.Remove(e)

	return val
}

func (this *MyStack) Top() int {

	e := this.list.Back()
	if e == nil {
		return 0
	}

	return e.Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.list.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}

func isValid(s string) bool {
	bytes := []byte(s)
	if len(bytes) == 0 || len(bytes)%2 != 0 {
		return false
	}

	hash := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]byte, 0)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '(' || bytes[i] == '[' || bytes[i] == '{' {
			stack = append(stack, hash[bytes[i]])
		} else {

			if len(stack) == 0 || stack[len(stack)-1] != bytes[i] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
