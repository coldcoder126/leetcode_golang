package main

import "fmt"

func main() {
	nums := [...]int{0, 0, 1}
	nums2 := nums[0:len(nums)]
	moveZeroes(nums2)
	fmt.Println(nums2)
}

func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}
