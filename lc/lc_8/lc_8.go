package main

import (
	"fmt"
	"math"
)

func main() {
	res := myAtoi("2147483646")
	fmt.Println(res)
}

func myAtoi(s string) int {
	res, sign, i, n := 0, 1, 0, len(s)
	//1. 去空格
	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}
	for ; i < n; i++ {
		if s[i] < 48 || s[i] > 57 {
			break
		}

		res = res*10 + int(s[i]-'0')
		if sign*res < math.MinInt32 {
			return math.MinInt32
		}
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return sign * res
}
