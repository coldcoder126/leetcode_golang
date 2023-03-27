package main

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {
	s := "tmmabcmt"
	// pwwkew
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
func lengthOfLongestSubstring(s string) int {
	arr := []rune(s)
	l := len(s)
	max_l := 0
	distance := 0
	mp := make(map[rune]int)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if v, ok := mp[arr[j]]; ok {
				if v >= i {
					i = v + 1
					distance = j - i + 1
				} else {
					distance++
					max_l = max(distance, max_l)
				}
			} else {
				distance++
				max_l = max(distance, max_l)
			}
			mp[arr[j]] = j
			if j == l-1 {
				return max_l
			}
		}
	}
	return max_l
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
