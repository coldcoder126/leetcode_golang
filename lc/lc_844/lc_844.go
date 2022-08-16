package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "ab##"
	str2 := "c#d#"
	fmt.Println(backspaceCompare(str1, str2))

}
func backspaceCompare(s string, t string) bool {
	// 两个分别处理，处理完再比较
	s1 := exec(s)
	s2 := exec(t)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func exec(str string) []string {
	strarr := strings.Split(str, "")

	pos := 0
	for i := 0; i < len(strarr); i++ {
		if strarr[i] != "#" {
			strarr[pos] = strarr[i]
			pos++
		} else {
			if pos > 0 {
				pos--
			}

		}
	}
	return strarr[:pos]
}

func getString(s string) string {
	bz := []rune{}
	for _, c := range s {
		if c != '#' {
			bz = append(bz, c)
		} else if len(bz) > 0 {
			bz = bz[:len(bz)-1]
		}
	}
	return string(bz)
}
