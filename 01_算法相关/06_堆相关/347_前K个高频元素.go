package main

import "container/heap"

// IHeap 构建小顶堆
type IHeap [][2]int

// Len 实现sort的interface方法 Len
func (h IHeap) Len() int {
	return len(h)
}

// Less 实现sort的interface方法 Less
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

// Swap 实现sort的interface方法 Swap
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 实现heap的内置两个方法 Push
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// Pop 实现heap的内置两个方法 Pop
func (h *IHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// 347_前K个高频元素: 无容量限制的小顶堆 + 内置heap方法
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		m[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, count := range m {
		heap.Push(h, [2]int{key, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
