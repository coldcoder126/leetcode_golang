package main

func findDisappearedNumbers(nums []int) []int {
	mark := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		mark[nums[i]] = 1
	}
	res := make([]int, 0)
	for i := 1; i < len(nums)+1; i++ {
		if mark[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}
