package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 维护一个单调递增队列 从的到小

type QueueInterface interface {
	Pop(val int)
	Push(val int)
	GetMax() int
}

type ListQueue struct {
	list *list.List
}

func (lq *ListQueue) Pop(val int) {
	if lq.list.Len() != 0 && lq.list.Front().Value.(int) == val {
		lq.list.Remove(lq.list.Front())
	}
}

func (lq *ListQueue) Push(val int) {
	for lq.list.Len() != 0 && lq.list.Back().Value.(int) < val {
		lq.list.Remove(lq.list.Back())
	}
	lq.list.PushBack(val)
}

func (lq *ListQueue) GetMax() int {
	if lq.list.Len() != 0 {
		return lq.list.Front().Value.(int)
	}
	return math.MinInt
}

func maxSlidingWindow(nums []int, k int) []int {

	q := &ListQueue{
		list: list.New(),
	}
	ans := make([]int, 0)

	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}

	ans = append(ans, q.GetMax())
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		ans = append(ans, q.GetMax())
	}

	return ans
}
