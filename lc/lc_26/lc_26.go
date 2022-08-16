package main

import "fmt"

func main() {
	nums := [...]int{1, 2, 3}
	nums2 := nums[0:len(nums)]
	res := removeDuplicates(nums2)
	fmt.Println(res)
	fmt.Println(nums2)
}

func removeDuplicates(nums []int) int {
	//思路，遍历并和下一个比较
	pre, next := 0, 0
	if len(nums) <= 1 {
		return len(nums)
	}
	for next < len(nums)-1 {
		next++
		if nums[next] == nums[pre] {
			continue
		} else {
			pre++
			nums[pre] = nums[next]
		}
	}

	return pre + 1
}
