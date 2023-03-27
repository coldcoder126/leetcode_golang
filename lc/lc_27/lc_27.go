package main

import "fmt"

func main() {
	nums := [...]int{3, 1, 3, 3, 2, 3, 4}
	nums2 := nums[:]
	res := removeElement2(nums2, 3)
	fmt.Println(res)
	fmt.Println(nums2)

}

func removeElement(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

func removeElement2(nums []int, val int) int {
	//两个指针
	low := 0
	high := 0
	l := len(nums)
	for high = 0; high < l; high++ {
		if nums[high] == val {
			continue
		}
		nums[low] = nums[high]
		low++
	}
	return low
}
