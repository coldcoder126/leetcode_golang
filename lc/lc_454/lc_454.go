package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	ans := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Println(ans)
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp1 := make(map[int]int)
	mp2 := make(map[int]int)

	n := len(nums1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := nums1[i] + nums2[j]
			mp1[sum]++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := nums3[i] + nums4[j]
			mp2[sum]++
		}
	}

	count := 0
	for k1, v1 := range mp1 {
		v2, ok := mp2[-k1]
		if ok {
			count += v1 * v2
			fmt.Println(v1, v2)
		}
	}
	return count
}
