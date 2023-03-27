package main

import "fmt"

func main() {
	nums := [...]int{-4, -1, 0, 3, 10}
	n2 := nums[:]
	res := sortedSquares(n2)
	fmt.Println(res)
}

// 从两端往中间找，边找边写入答案
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	idx := right
	for idx >= 0 {
		l := nums[left] * nums[left]
		r := nums[right] * nums[right]
		if l >= r {
			res[idx] = l
			left++
		} else {
			res[idx] = r
			right--
		}
		idx--
	}
	return res
}
