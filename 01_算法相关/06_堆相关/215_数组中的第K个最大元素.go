package main

import "container/heap"

// 215_数组中的第K个最大元素: 固定容量小顶堆
func findKthLargest(nums []int, k int) int {
	// 初始化小顶堆
	minHeap := newMinHeap(k)
	// 数据入小顶堆
	for _, num := range nums {
		minHeap.add(num)
	}
	num, _ := minHeap.peek()
	return num
}

type mHeap []int

// Less 实现sort的interface方法 Less
func (h mHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap 实现sort的interface方法 Swap
func (h mHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Len 实现sort的interface方法 Len
func (h mHeap) Len() int {
	return len(h)
}

// Push 实现heap的interface方法 Push
func (h *mHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 实现heap的interface方法 Pop
func (h *mHeap) Pop() interface{} {
	temp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return temp
}

// 215_数组中的第K个最大元素: 基于内置heap实现
func findKthLargest2(nums []int, k int) int {
	// 初始化小顶堆
	minHeap := &mHeap{}
	heap.Init(minHeap)
	// 数据入小顶堆
	for _, num := range nums {
		heap.Push(minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	return (*minHeap)[0]
}
