package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(arr)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {
	mp := make(map[int]int)
	mp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		for j := len(mp); j > 0; j-- {
			if val, ok := mp[j]; ok {
				if nums[i] <= val {
					if j == 1 {
						mp[j] = nums[i]
					}
					continue
				} else {
					mp[j+1] = nums[i]
					break
				}
			}
		}
	}
	return len(mp)
}
