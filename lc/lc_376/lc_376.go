package main

func wiggleMaxLength(nums []int) int {
	l := len(nums)
	count := 0
	gateHigh := true
	gateLow := true
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] > nums[i-1] && gateHigh {
			count++
			gateHigh = false
			gateLow = true
			continue
		}
		if i > 0 && nums[i] < nums[i-1] && gateLow {
			count++
			gateHigh = true
			gateLow = false
			continue
		}
	}
	return count + 1
}
