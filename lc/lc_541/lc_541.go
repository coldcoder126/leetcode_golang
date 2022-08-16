package main

import "fmt"

func main() {
	ans := reverseStr("abcdefghijklmnopq", 3)
	fmt.Println(ans)
}
func reverseStr(s string, k int) string {
	arr := []byte(s)
	left, right := 0, 0
	length := len(s) - 1
	step := 1
	for right < length {
		right++
		step++
		if right > 2*k-1 {
			left++
		}
		if (right+1)%(2*k) == 0 {
			reverse(arr, left, left+k-1)
			step = 0
		} else if right == length {
			if step >= k {
				reverse(arr, right-step+1, right-step+k)
			} else {
				reverse(arr, right-step+1, right)
			}
		}

	}
	return string(arr)
}

func reverse(str []byte, start, end int) {
	for start < end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}
}
