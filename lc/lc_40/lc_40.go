package main

import (
	"fmt"
	"sort"
)

func main() {
	abc := []int{10, 1, 2, 7, 6, 1, 5}
	sort.Ints(abc)
	x := combinationSum2(abc, 8)
	fmt.Println(x)
}

func combinationSum2(candidates []int, target int) [][]int {
	// 排序
	sort.Ints(candidates)
	flag := make([]int, len(candidates))
	var res []int
	var ans [][]int
	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			temp := make([]int, len(res))
			copy(temp, res)
			ans = append(ans, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > 0 && flag[i-1] == 0 && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] <= target {
				flag[i] = 1
				res = append(res, candidates[i])
				dfs(i+1, target-candidates[i])
				res = res[:len(res)-1]
				flag[i] = 0
			}
		}
	}
	dfs(0, target)
	return ans
}
