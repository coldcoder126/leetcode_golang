package main

import "fmt"

func main() {
	arr := []int{1}
	nums := arr[:]
	res := searchRange(nums, 1)
	fmt.Println(res)
}

// 思路：二分法找到目标位置，再向前向后遍历
func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	pos, begin, end := -1, -1, -1

	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			pos = mid
			break
		}
	}
	if pos == -1 {
		return []int{-1, -1}
	}
	for i := pos; i <= len(nums)-1; i++ {
		if nums[i] == nums[pos] {
			end = i
		} else {
			break
		}
	}

	for i := pos; i >= 0; i-- {
		if nums[i] == nums[pos] {
			begin = i
		} else {
			break
		}
	}
	return []int{begin, end}
}

//优化：查找左右边界的时候也使用二分法查找
