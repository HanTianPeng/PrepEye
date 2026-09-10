package main

type MyQueue struct {
	Stack []int
	Size  int
}

func Constructor() MyQueue {
	return MyQueue{
		Stack: make([]int, 0),
		Size:  0,
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
	this.Size++
}

func (this *MyQueue) Pop() int {
	if this.Size == 0 {
		return -1
	}
	temp := this.Stack[0]
	this.Stack = this.Stack[1:]
	this.Size--
	return temp
}

func (this *MyQueue) Peek() int {
	return this.Stack[0]
}

func (this *MyQueue) Empty() bool {
	return this.Size == 0
}
