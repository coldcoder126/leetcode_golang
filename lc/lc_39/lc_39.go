package main

import "fmt"

func main() {
	input := []int{2, 3, 6, 7}

	a := combinationSum(input, 7)
	fmt.Println(a)
}

func combinationSum(candidates []int, target int) [][]int {
	var res []int
	var ans [][]int

	var dfs func(input []int, start int, target int)

	dfs = func(input []int, start int, target int) {
		if target == 0 {
			temp := make([]int, len(res))
			copy(temp, res)
			ans = append(ans, temp)
			return
		}

		for i := start; i < len(input); i++ {
			if input[i] <= target {
				res = append(res, input[i])
				dfs(input, i, target-input[i])
				if len(res) > 0 {
					res = (res)[:len(res)-1]
				}
			}
		}
	}

	dfs(candidates, 0, target)
	return ans
}


