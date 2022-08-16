package main

import "fmt"

func main() {
	ans := isAnagram("rat", "car")
	fmt.Println(ans)
}

func isAnagram(s string, t string) bool {
	length := len(s)
	if length != len(t) {
		return false
	}

	arr1 := []byte(s)
	arr2 := []byte(t)
	m := make(map[byte]int)
	for i := 0; i < length; i++ {
		m[arr1[i]]++
		m[arr2[i]]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
