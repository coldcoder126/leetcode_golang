package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := reverse(1563847412)
	fmt.Println(x)
}
func reverse(x int) int {
	//第一步，将正负号提取出来
	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}
	max := math.MaxInt32
	arr := make([]int, 0)
	//第二步，入队
	for i := x; i > 0; i = int(i / 10) {
		t := i % 10
		arr = append(arr, t)
	}
	//第三步，去除开头的0
	first_is_zero := true
	s := ""
	for _, val := range arr {
		if val > 0 && first_is_zero {
			first_is_zero = false
		}
		if !first_is_zero {
			s += strconv.Itoa(val)
		}
	}
	res, _ := strconv.Atoi(s)
	res = res * flag
	if res > max || res < -max {
		return 0
	} else {
		return res
	}
}
