package main

import "fmt"

func main() {
	m := []int{3, 4}
	n := []int{1, 2}
	res := findMedianSortedArrays(m, n)
	fmt.Println(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l3 := (l1+l2)/2 + 1
	j, k := 0, 0
	t := make([]int, 2)
	for j+k < l3 {
		if j < l1 {
			if k < l2 && nums2[k] < nums1[j] {
				t[0] = t[1]
				t[1] = nums2[k]
				k++
			} else {
				t[0] = t[1]
				t[1] = nums1[j]
				j++
			}
		} else {
			t[0] = t[1]
			t[1] = nums2[k]
			k++
		}
	}
	if (l1+l2)%2 == 0 {
		return float64(t[0]+t[1]) / 2.0
	} else {
		return float64(t[1])
	}
}
