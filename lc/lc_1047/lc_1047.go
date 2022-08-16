package main

import "fmt"

func main() {
	s := "abbaca"
	res := removeDuplicates(s)
	fmt.Println(res)
}

func removeDuplicates(s string) string {
	// 消消乐
	arr := []byte(s)
	var temp []byte
	for _, i := range arr {
		l := len(temp)
		if l > 0 && i == temp[l-1] {
			temp = temp[:l-1]
		} else {
			temp = append(temp, i)
		}
	}
	return string(temp)
}
