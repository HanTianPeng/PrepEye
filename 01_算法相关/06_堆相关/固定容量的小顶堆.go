package main

import "fmt"

// MinHeap 实现一个小顶堆: 任意父节点<=子节点
type MinHeap struct {
	data []int
	size int
}

func newMinHeap(size int) *MinHeap {
	return &MinHeap{data: []int{}, size: size}
}

// swap: 交换元素
func (hp *MinHeap) swap(i, j int) {
	hp.data[j], hp.data[i] = hp.data[i], hp.data[j]
}

// len: 返回实际堆大小
func (hp *MinHeap) len() int {
	return len(hp.data)
}

func (hp *MinHeap) less(parent, child int) bool {
	return hp.data[parent] < hp.data[child]
}

// peek: 获取堆顶元素
func (hp *MinHeap) peek() (int, bool) {
	if hp.len() == 0 {
		return 0, false
	}
	return hp.data[0], true
}

func (hp *MinHeap) print() {
	fmt.Print(hp.data)
}

// up: 上浮
func (hp *MinHeap) up(j int) {
	for {
		parent := (j - 1) / 2
		// 新元素比父节点大,这不用上浮
		if parent == j || hp.less(parent, j) {
			break
		}
		// 如果新元素比父节点小,这需要上浮
		hp.swap(parent, j)
		j = parent
	}
}

// down: 下沉
func (hp *MinHeap) down(i int) {
	l := hp.len()
	parent := i
	for {
		left := parent*2 + 1
		// 如果左指针超过了数组的最大长度,直接返回
		if left >= l || left < 0 {
			break
		}
		// 先取出左节点,右节点可能不存在
		j := left
		if right := j + 1; right < l && hp.less(right, left) {
			j = right
		}
		// 等于当前，证明找到了合适的位置
		if hp.less(parent, j) {
			break
		}
		// 没有找到合适位置,交换
		hp.swap(parent, j)
		parent = j
	}
}

// add: 只保留最大的k个元素，堆顶为最大里面最小的元素
func (hp *MinHeap) add(num int) {
	// 固定容量不能小于等于0
	if hp.size <= 0 {
		return
	}
	// 堆还没有满
	if hp.len() < hp.size {
		hp.data = append(hp.data, num)
		hp.up(hp.len() - 1)
		return
	}
	// 堆满了，堆顶是当前TopK中最小的，当前元素小于堆顶，直接抛弃
	if num < hp.data[0] {
		return
	}
	// 堆满了,堆顶是当前TopK中最小的，当前元素大于堆顶，则需要替换堆顶并下沉
	hp.data[0] = num
	hp.down(0)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 4

	heap := newMinHeap(k)

	for _, num := range nums {
		heap.add(num)
	}
	heap.print() // 4
}
