package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 2}
	res := findDuplicate(arr)
	fmt.Println(res)
}

func findDuplicate(nums []int) int {
	//快慢指针
	slow := 0
	fast := 0
	for i := 0; i < len(nums); i++ {
		slow = nums[slow]
		fast = nums[nums[fast]]
		// 相遇时 fast一定是值而不是索引
		if slow == fast {
			p := 0
			for p != slow {
				p = nums[p]
				slow = nums[slow]
			}
			return slow
		}
	}
	return -1
}
