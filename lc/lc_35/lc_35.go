package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6, 9}
	nums := arr[:]
	res := searchInsert2(nums, 8)
	fmt.Println(res)
}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
//如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var pos int
	for left <= right {
		mid := (left + right) / 2
		if left == mid {
			pos = left
		}
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	if nums[pos] > target {
		return pos
	} else {
		return pos + 1
	}
}

func searchInsert2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	// 分别处理如下四种情况
	// 目标值在数组所有元素之前  [0, -1]
	// 目标值等于数组中某一个元素  return middle;
	// 目标值插入数组中的位置 [left, right]，return  right + 1 【这个比较难理解】
	// 目标值在数组所有元素之后的情况 [left, right]， 因为是右闭区间，所以 return right + 1
	return r + 1
}
