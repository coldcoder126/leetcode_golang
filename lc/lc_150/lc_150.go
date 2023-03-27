package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	res := evalRPN(tokens)
	fmt.Println(res)
}

/**
执行用时：4 ms, 在所有 Go 提交中击败了82.19%的用户
内存消耗：3.9 MB, 在所有 Go 提交中击败了87.15%的用户
通过测试用例：21 / 21
*/

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	var x1, x2, x3 int
	for _, c := range tokens {
		switch c {

		case "+":
			x1 = stack[len(stack)-1]
			x2 = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			x3 = x1 + x2
		case "-":
			x1 = stack[len(stack)-1]
			x2 = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			x3 = x2 - x1
		case "*":
			x1 = stack[len(stack)-1]
			x2 = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			x3 = x1 * x2
		case "/":
			x1 = stack[len(stack)-1]
			x2 = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			x3 = x2 / x1
		default:
			x3, _ = strconv.Atoi(c)
		}
		stack = append(stack, x3)
	}
	return stack[0]
}
