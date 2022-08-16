package main

import "fmt"

func main() {
	x := 2147395600
	res := mySqrt(x)
	fmt.Println(res)
}

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		val := (left + right) / 2
		if val*val < x {
			left = val + 1
		} else if val*val > x {
			right = val - 1
		} else {
			return val
		}
	}
	return right
}
