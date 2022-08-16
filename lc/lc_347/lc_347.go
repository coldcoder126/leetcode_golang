package main

import (
	"sort"
)

func main() {
	//nums := []int{1, 2}
	//ans := topKFrequent(nums, 2)
	//fmt.Println(ans)

}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, i := range nums {
		mp[i]++
	}

	// map按值排序
	mp2 := make(map[int][]int)
	arr := make([]int, 0)
	for key, count := range mp {
		_, ok := mp2[count]
		if ok {
			mp2[count] = append(mp2[count], key)
		} else {
			mp2[count] = []int{key}
		}
	}

	for K, _ := range mp2 {
		arr = append(arr, K)
	}

	res := make([]int, 0)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for i := range arr[:] {
		val, _ := mp2[arr[i]]
		res = append(res, val...)
	}

	return res[:k]
}

func topKFrequent2(nums []int, k int) []int {
	ans := []int{}
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans, func(a, b int) bool {
		return map_num[ans[a]] > map_num[ans[b]]
	})
	return ans[:k]
}
