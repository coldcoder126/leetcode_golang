package main

import "fmt"

func main() {
	res := numTrees(3)
	fmt.Println(res)
}

func numTrees(n int) int {
	mp := make([]int, n+1)
	mp[0] = 1
	mp[1] = 1
	for i := 2; i <= n; i++ {
		count := 0
		for j := 0; j < i; j++ {
			count += mp[j] * mp[i-j-1]
		}
		mp[i] = count
	}
	return mp[n]
}
