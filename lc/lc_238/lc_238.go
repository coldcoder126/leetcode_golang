package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)

	left[0] = nums[0]
	right[l-1] = nums[l-1]

	for i := 1; i < l; i++ {
		left[i] = nums[i] * left[i-1]
		right[i] = nums[l-1-i] * right[l-i]
	}

	res := make([]int, l)
	res[0] = right[1]
	res[l-1] = left[l-2]
	for i := 1; i < l-1; i++ {
		res[i] = left[i-1] * right[i+1]
	}
	return res

}
