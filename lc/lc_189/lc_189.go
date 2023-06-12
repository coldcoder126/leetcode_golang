package main

import "fmt"

//21:22
func main() {
	arr := []int{1}
	rotate(arr, 2)
	fmt.Println(arr)
}

// 用反转的方法
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
