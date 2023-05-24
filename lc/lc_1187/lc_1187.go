package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{1, 5, 3, 6, 7}
	arr2 := []int{1, 6, 3, 3}
	res := makeArrayIncreasing(arr1, arr2)
	fmt.Println(res)
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)

	l1, l2 := len(arr1), len(arr2)
	j := l2 - 1
	//倒序遍历
	for i := l1 - 1; i > 0; i-- {
		if arr1[i-1] >= arr1[i] {
			//如果遇到非严格递减，则从arr2中找一个符合要求最大的替代
			for ; j > 0 && arr2[j] >= arr1[i+1]; j-- {
			}
			arr1[i] = arr2[j]
			i++
		}
	}
	return 1
}
