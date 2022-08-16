package main

import "fmt"

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	str := "abcdefg"
	s1 := []byte(str)
	fmt.Println(s1)
	reverseString(s1)
	fmt.Println(string(s1))
}
