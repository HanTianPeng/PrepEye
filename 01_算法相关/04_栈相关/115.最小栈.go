package main

import "math"

type MinStack struct {
	Data []pair
}

type pair struct {
	val    int
	preMin int
}

func StackConstructor() MinStack {
	item := pair{
		val:    0,
		preMin: math.MaxInt,
	}
	return MinStack{
		Data: []pair{item},
	}
}

func (this *MinStack) Push(val int) {
	this.Data = append(this.Data, pair{
		val:    val,
		preMin: min(this.GetMin(), val),
	})
}

func (this *MinStack) Pop() {
	this.Data = this.Data[:len(this.Data)-1]
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1].val
}

func (this *MinStack) GetMin() int {
	return this.Data[len(this.Data)-1].preMin
}
