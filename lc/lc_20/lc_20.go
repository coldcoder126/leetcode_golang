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

func isValid2(s string) bool {
	mp := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]byte, 0)
	for i, _ := range s {
		if len(stack) > 0 {
			x := stack[len(stack)-1]
			if mp[x] == s[i] {
				stack = stack[:len(stack)-1] //出栈
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
