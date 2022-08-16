package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	nums := arr[:]
	res := search(nums, 9)
	fmt.Println(res)
}

// 给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
//写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

// 第一种写法[left,right]左闭右闭
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 第二种写法[left,right)左闭右开
func search2(nums []int, target int) int {
	high := len(nums)
	low := 0
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
