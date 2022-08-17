package lc_55

func canJump(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums); i++ {
		maxPos = max(i+nums[i], maxPos)
		if i < len(nums)-1 && maxPos <= i {
			return false
		}
		if maxPos >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
