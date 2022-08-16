package main

import "fmt"

func main() {
	ans := reverseLeftWords2("lrloseumgh", 6)
	fmt.Println(ans)
}
func reverseLeftWords(s string, n int) string {
	arr := []byte(s)
	arr = append(arr, arr[:n]...)
	return string(arr[n:])
}

// 思路2：不用额外空间
func reverseLeftWords2(s string, n int) string {
	arr := []byte(s)
	s1 := arr[:n]
	s2 := arr[n:]
	reverse(s1)
	reverse(s2)
	reverse(arr)
	return string(arr)
}

func reverse(arr []byte) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
