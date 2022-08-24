package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	stack.Pop()
	x := stack.GetMin()
	fmt.Println(x)
}

type MinStack struct {
	size int
	val  []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{size: 0, val: make([]int, 0), min: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	var minVal int
	if this.size == 0 {
		minVal = val
	} else {
		minVal = min(this.min[this.size-1], val)
	}
	this.size++
	this.val = append(this.val, val)
	this.min = append(this.min, minVal)
}

func (this *MinStack) Pop() {
	this.size--
	this.val = this.val[:this.size]
	this.min = this.min[:this.size]

}

func (this *MinStack) Top() int {
	return this.val[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
