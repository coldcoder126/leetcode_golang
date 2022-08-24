package main

func majorityElement(nums []int) int {
	mp := make(map[int]int)
	l := len(nums)
	half := l / 2
	for i := 0; i < l; i++ {
		mp[nums[i]]++
		if i >= half {
			if mp[nums[i]] > half {
				return nums[i]
			}
		}
	}
	return 0
}

//进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题
// 摩尔投票法：不同元素之间相互抵消，极端情况是主体元素A和其他所有元素相互抵消，最后剩下的一定是主体元素

func majorityElement2(nums []int) int {
	l := len(nums)
	count := 1
	candidate := nums[0]
	for i := 1; i < l; i++ {
		if count == 0 {
			candidate = nums[i]
			count++
			continue
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
