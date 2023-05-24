package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	//target := 9
	//result := twoSum(nums, target)
	//fmt.Println(result)
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}

// 两遍遍历
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		if k, ok := mp[target-nums[i]]; ok {
			return []int{k, i}
		} else {
			mp[nums[i]] = i
		}
	}
	return nil
}
