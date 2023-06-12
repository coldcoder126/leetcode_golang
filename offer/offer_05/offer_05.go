package main

import "fmt"

func main() {
	res := replaceSpace(" hello  world ")
	fmt.Println(res)
}

func replaceSpace(s string) string {
	arr := []byte(s)
	l1 := len(s) - 1
	count := 0
	//遍历，记录空格的个数
	for _, char := range arr {
		if char == ' ' {
			count++
		}
	}
	arr = append(arr, make([]byte, count*2)...)
	l2 := l1 + count*2

	for l1 >= 0 {
		if arr[l1] == ' ' {
			arr[l2] = '0'
			arr[l2-1] = '2'
			arr[l2-2] = '%'
			l2 = l2 - 3
		} else {
			arr[l2] = arr[l1]
			l2--
		}
		l1--
	}
	return string(arr)
}
