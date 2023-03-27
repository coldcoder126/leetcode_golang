package main

import "fmt"

func main() {
	a := 12
	flag := isPalindrome(a)
	fmt.Println(flag)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//1.将x分解为数组
	arr := make([]int, 0)
	for i := x; i > 0; i = int(i / 10) {
		arr = append(arr, i%10)
	}
	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-1-i] {
			return false
		}
	}
	return true
}
