package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	res := subsets(arr)
	fmt.Println(res)

}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0)

	var dfs func(begin int)
	dfs = func(begin int) {
		t := make([]int, len(temp))
		copy(t, temp)
		ans = append(ans, t)

		for i := begin; i < len(nums); i++ {
			temp = append(temp, nums[i])
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return ans
}
