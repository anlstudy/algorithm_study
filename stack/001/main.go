package main

import (
	"math"
)

//设计一个支持push，pop，top，getMin操作，并且能在常数时间内检索最小元素的栈
//push（x）
//pop（）
//top（）
//getMin（）

type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{}, math.MinInt64,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.data) == 0 {
		this.min = x
	} else {
		if this.min > x {
			this.min = x
		}
	}

}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	min := math.MaxInt64
	for _, v := range this.data {
		if v < min {
			min = v
		}
	}
	this.min = min

}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]

}

func (this *MinStack) GetMin() int {
	return this.min
}
