package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	ans := topKFrequent(nums, 2)
	fmt.Println(ans)
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

// 使用堆来时刻保存频率最高的k个，堆的构造为二维数组 a[i][0]是数值，a[i][1]是频率
//为什么使用大顶堆？因为小顶堆的Pop操作会将目前最小的弹出，最后保留的k个里，是最大的k个

/*
*
执行用时：12 ms, 在所有 Go 提交中击败了77.31%的用户
内存消耗：5.2 MB, 在所有 Go 提交中击败了65.46%的用户
通过测试用例：21 / 21
*/
type fHeap [][2]int

func (h fHeap) Len() int {
	return len(h)
}

func (h fHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h fHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *fHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *fHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func topKFrequent3(nums []int, k int) []int {
	h := &fHeap{}
	heap.Init(h)

	mp := make(map[int]int)

	for _, item := range nums {
		mp[item]++
	}

	for key, val := range mp {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}
