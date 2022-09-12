package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

func main() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//fmt.Println(maxSlidingWindow(nums, 3))

	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
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

type elem struct {
	num, count int
}

type priQueue []elem

func (p priQueue) Len() int {
	return len(p)
}

func (p priQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p priQueue) Less(i, j int) bool {
	return p[i].count < p[j].count
}

func (p *priQueue) Push(x interface{}) {
	*p = append(*p, x.(elem))
}

func (p *priQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, k, k)
	counts := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}

	h := &priQueue{}
	heap.Init(h)
	for num, count := range counts {
		heap.Push(h, elem{
			num:   num,
			count: count,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	for i := h.Len() - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(elem).num
	}

	return ans
}
