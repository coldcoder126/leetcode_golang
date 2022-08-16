package main

func main() {
	combinationSum3(3, 7)
}

func combinationSum3(k int, n int) [][]int {
	var res []int
	var ans [][]int
	dfs(1, n, k, &res, &ans)
	return ans
}

func dfs(start, n, k int, res *[]int, ans *[][]int) {
	if k == 0 && n == 0 {
		temp := make([]int, len(*res))
		copy(temp, *res)
		*ans = append(*ans, temp)
		return
	}
	for i := start; i < 9; i++ {
		if i <= n {
			*res = append(*res, i)
			dfs(i+1, n-i, k-1, res, ans)
			*res = (*res)[:len(*res)-1]
		} else {
			return
		}
	}
}
