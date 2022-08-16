package main

import "fmt"

func main() {
	res := generateMatrix(3)
	fmt.Println(res)
}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	ans[0][0] = 1
	x, y := 0, 0
	layer := 1
	val := 1
	for val < n*n {
		for y < n-layer {
			val++
			y++
			ans[x][y] = val
		}

		for x < n-layer {
			val++
			x++
			ans[x][y] = val
		}

		for y > layer-1 {
			val++
			y--
			ans[x][y] = val
		}
		for x > layer {
			val++
			x--
			ans[x][y] = val
		}
		layer++
	}
	return ans
}
