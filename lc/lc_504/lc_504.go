package main

import (
	"fmt"
	"strconv"
)

func main() {

	res := convertToBase7(100)
	fmt.Println(res)
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	flag := true
	x := num
	if num < 0 {
		x = -num
		flag = false
	}

	for x != 0 {
		i := x % 7
		res = strconv.Itoa(i) + res
		x /= 7
	}

	if flag {
		return res
	} else {
		return "-" + res
	}
}
