package main

import "fmt"

func main() {
	s := "([)]"
	ans := isValid(s)
	fmt.Println(ans)
}

func isValid(s string) bool {
	arr := []byte(s)
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var temp []byte
	for _, i := range arr {
		top := len(temp)
		val, ok := mp[i]
		if top > 0 && ok && temp[top-1] == val {
			temp = temp[:top-1]
		} else {
			temp = append(temp, i)
		}
	}
	return len(temp) == 0
}
