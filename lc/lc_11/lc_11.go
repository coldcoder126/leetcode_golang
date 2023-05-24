package main

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(arr)
	fmt.Println(res)
}

func maxArea(height []int) int {
	//双指针
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area = max(area, w*h)
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
