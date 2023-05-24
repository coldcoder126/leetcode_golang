package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	s := "112"

	//arr := []byte(s)
	//arr = append(arr[0:1], arr[2:]...)
	//fmt.Println(string(arr))
	ans := removeKdigits2(s, 1)
	fmt.Println(ans)
}

// dfs法
func removeKdigits(num string, k int) string {
	arr := []byte(num)
	l := len(arr) - k
	if l == 0 {
		return "0"
	}

	minVal := math.MaxInt32

	res := make([]byte, 0)
	var dfs func(begin int)
	dfs = func(begin int) {
		if len(res) == l {
			//结束
			//fmt.Println(string(res))
			val := removeFrontZero(res)
			if x, err := strconv.Atoi(string(val)); err == nil {
				minVal = min(minVal, x)
			}
		}

		//不结束，继续
		for i := begin; i < len(arr); i++ {
			res = append(res, arr[i])
			dfs(i + 1)
			res = res[:len(res)-1]
		}
	}
	dfs(0)
	return strconv.Itoa(minVal)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func removeFrontZero(arr []byte) []byte {
	notZeroIdx := len(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != byte('0') {
			notZeroIdx = i
			break
		}
	}
	return arr[notZeroIdx:]
}

// 贪心法，从左到右，删除最左边最大的，还是会超时
func removeKdigits2(num string, k int) string {
	arr := []byte(num)
	l := len(arr) - k
	if l == 0 {
		return "0"
	}
	for i := 0; i < k; i++ {
		foundFlag := false
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr = append(arr[:j-1], arr[j:]...)
				foundFlag = true
				break
			}
		}
		if !foundFlag {
			arr = arr[:len(arr)-1]
		}
	}
	ans := removeFrontZero(arr)
	if len(ans) == 0 {
		return "0"
	} else {
		return string(ans)
	}
}

// 贪心+单调栈 才不超时
func removeKdigits3(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
