package main

import "fmt"

// MaxHeap 实现一个大顶堆: 父节点 >= 子节点;
type MaxHeap struct {
	data []int
	size int
}

// newMaxHeap 初始化
func newMaxHeap(size int) *MaxHeap {
	return &MaxHeap{data: []int{}, size: size}
}

// swap 交换位置函数
func (hp *MaxHeap) swap(i, j int) {
	hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}

// len 堆实际大小
func (hp *MaxHeap) len() int {
	return len(hp.data)
}

// peek 获取堆顶元素
func (hp *MaxHeap) peek() (int, bool) {
	if hp.len() == 0 {
		return 0, false
	}
	return hp.data[0], true
}

func (hp *MaxHeap) print() {
	fmt.Print(hp.data)
}

// add 添加新元素
func (hp *MaxHeap) add(num int) {
	// 当前堆没有满
	if hp.len() < hp.size {
		hp.data = append(hp.data, num)
		hp.up(hp.len() - 1)
		return
	}
	// 堆满了，新元素大于堆顶元素，堆顶是最小的最大元素
	n, _ := hp.peek()
	if num > n {
		return
	}
	// 堆满了，堆顶是最小元素中的最大元素,当前元素小于堆顶，则需要替换并下沉
	hp.data[0] = num
	hp.down(0)
}

// up 上浮
func (hp *MaxHeap) up(j int) {
	for {
		parent := (j - 1) / 2
		// 新元素比父节点小，不需要上浮
		if parent == j || hp.data[parent] >= hp.data[j] {
			break
		}
		// 如果新元素比父节点大,则需要上浮
		hp.swap(parent, j)
		j = parent
	}
}

// down 下沉
func (hp *MaxHeap) down(i int) {
	l := hp.len()
	largest := i
	for {
		left, right := 2*i+1, 2*i+2
		if left < l && hp.data[largest] < hp.data[left] {
			largest = left
		}
		if right < l && hp.data[largest] < hp.data[right] {
			largest = right
		}
		if largest == i {
			break
		}
		hp.swap(largest, i)
		i = largest
	}
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4, 7}
	k := 4

	heap := newMaxHeap(k)

	for _, num := range nums {
		heap.add(num)
	}
	heap.print() // 3
}
