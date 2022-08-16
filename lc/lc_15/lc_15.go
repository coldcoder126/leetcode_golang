package main

import (
	"fmt"
	"sort"
)

func main() {
	// [-1,0,1,2,-1,-4,-2,-3,3,0,4]
	nums := []int{0, 0, 0, 0, 0, 0}
	ans := threeSum(nums)
	fmt.Println(ans)
}

// 双指针，固定首个，移动后面两个指针
func threeSum(nums []int) [][]int {
	l := len(nums)
	var res [][]int
	if l < 3 {
		return res
	}

	// 1. 将数组排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	left, right := 1, l-1
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1
		right = l - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if sum > 0 {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			}
		}

	}

	return res
}
