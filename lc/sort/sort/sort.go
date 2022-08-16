package main

import "fmt"

func main() {
	arr := []int{9, 7, 2, 4, 1, 7, 1, 44, 2, 63, 2, 643, 42, 9}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 快排

func Partition(nums []int, low, high int) int {
	temp := nums[low]
	pivot := nums[low]
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= pivot {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = temp
	return low
}

func quick_sort(nums []int, low, high int) {
	if low < high {
		pivot := Partition(nums, low, high)
		quick_sort(nums, low, pivot-1)
		quick_sort(nums, pivot+1, high)
	}
}
