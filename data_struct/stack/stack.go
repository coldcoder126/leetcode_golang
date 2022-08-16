package main

import (
	"errors"
	"fmt"
)

// ---------数组实现----------
func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.val)
	fmt.Println(s.Pop())
	fmt.Println(s.val)
	fmt.Println(s.Peek())
	fmt.Println("栈的长度：", s.size)
	//fmt.Println("栈是否为空：", s.Empty())

}

type ArrStack struct {
	size int
	val  []int
}

func Constructor() ArrStack {
	return ArrStack{}
}

func (s *ArrStack) Push(x int) {
	s.size++
	s.val = append(s.val, x)
}

func (s *ArrStack) Pop() (int, error) {
	if s.size == 0 {
		return -1, errors.New("stack is empty")
	}
	res := s.val[s.size-1]
	s.size--
	s.val = s.val[:s.size]
	return res, nil
}

func (s *ArrStack) Peek() (int, error) {
	if s.size == 0 {
		return -1, errors.New("stack is empty")
	}
	return s.val[s.size-1], nil
}

func (s *ArrStack) Size() int {
	return s.size
}
