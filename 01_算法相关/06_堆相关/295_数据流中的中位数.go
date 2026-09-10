package main

import "container/heap"

type IntHeap []int

// minHeap1 小顶堆: 存较大一半元素: 堆顶元素 [4, 5, 6, 7]
type minHeap1 struct {
	IntHeap
}

// maxHeap1 大顶堆: 存较小一半元素: 堆顶元素 [3, 2, 1]
type maxHeap1 struct {
	IntHeap
}

// Len 长度计算 /*---------IntHeap通用函数----------*/
func (ih IntHeap) Len() int {
	return len(ih)
}

// Swap 交换
func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

// Push 追加
func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

// Pop 删除
func (ih *IntHeap) Pop() interface{} {
	l := len(*ih)
	temp := (*ih)[l-1]
	*ih = (*ih)[:l-1]
	return temp
}

func (mh minHeap1) Less(i, j int) bool {
	return mh.IntHeap[i] < mh.IntHeap[j]
}

func (mh maxHeap1) Less(i, j int) bool {
	return mh.IntHeap[i] > mh.IntHeap[j]
}

type MedianFinder struct {
	left  *minHeap1
	right *maxHeap1
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &minHeap1{},
		right: &maxHeap1{},
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == mf.right.Len() {
		heap.Push(mf.right, num)
		val := heap.Pop(mf.right).(int)
		heap.Push(mf.left, val)
	} else {
		heap.Push(mf.left, num)
		val := heap.Pop(mf.left).(int)
		heap.Push(mf.right, val)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() == mf.right.Len() {
		return float64(mf.left.IntHeap[0]+mf.right.IntHeap[0]) / 2.0
	}
	return float64(mf.left.IntHeap[0])
}

func main() {

}
