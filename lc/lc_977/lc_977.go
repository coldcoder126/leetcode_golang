package main

import (
	"fmt"
)

func main() {
	nums := [...]int{-7, -3, 2, 3, 11}
	n2 := nums[:]
	res := sortedSquares(n2)
	fmt.Println(res)
}

func sortedSquares(nums []int) []int {
	// 双指针，从两头开始往中间找
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	idx := len(nums) - 1
	for left <= right {
		if nums[left]*nums[left] <= nums[right]*nums[right] {
			res[idx] = nums[right] * nums[right]
			right--
		} else {
			res[idx] = nums[left] * nums[left]
			left++
		}
		idx--
	}
	return res
}
