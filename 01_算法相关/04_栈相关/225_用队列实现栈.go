package main

type MyStack struct {
	Queue []int
	Size  int
}

func Constructor() MyStack {
	return MyStack{
		Queue: make([]int, 0),
		Size:  0,
	}
}

func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
	this.Size++
}

func (this *MyStack) Pop() int {
	if this.Size == 0 {
		return -1
	}
	temp := this.Queue[this.Size-1]
	this.Queue = this.Queue[:this.Size-1]
	this.Size--
	return temp
}

func (this *MyStack) Top() int {
	return this.Queue[this.Size-1]
}

func (this *MyStack) Empty() bool {
	return this.Size == 0
}
