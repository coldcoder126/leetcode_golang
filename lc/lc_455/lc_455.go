package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	// 排序
	i, j := 0, 0 // i：孩子 j:饼干
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}
