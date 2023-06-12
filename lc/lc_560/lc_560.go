package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, -1, 1, -1}
	res := subarraySum(arr, 0)
	fmt.Println(res)

}

func subarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		//遍历[0,start]范围内所有可能的组合
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}
