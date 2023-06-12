package main

var res [][]int
var temp []int

func main() {
	combine(4, 2)
}
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	arr := make([]int, 0)
	var dfs func(begin int)
	dfs = func(begin int) {
		//终止条件
		if len(arr) == k {
			t := make([]int, k)
			copy(t, arr)
			ans = append(ans, t)
			return
		}

		for i := begin; i <= n; i++ {
			arr = append(arr, i)
			dfs(i + 1)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(1)
	return ans

}
