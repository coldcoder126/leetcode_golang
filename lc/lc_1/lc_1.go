package main

// 两遍遍历
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		mp[i] = nums[i]
	}

	for i := range nums {
		k, v := mp[i]
		if v {
			return []int{i, k}
		}
	}
	return nil
}

// 一遍遍历
func twoSum1(nums []int, target int) []int {
	mp := make(map[int]int)
	l := len(nums)

	for i := 0; i < l; i++ {
		k, v := mp[target-nums[i]]
		if v {
			return []int{i, k}
		} else {
			mp[nums[i]] = i
		}
	}
	return nil
}
