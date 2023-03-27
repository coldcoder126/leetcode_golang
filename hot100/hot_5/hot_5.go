package main

import "fmt"

func main() {
	s := "a"
	//aacabdkacaa,ccc,bb,tattarrattat,abac
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	// 有两种情况：aba式和abba式
	l := len(s)
	arr := []byte(s)
	pos1 := 0
	pos2 := 0
	arm1 := 0
	arm2 := 0

	if l == 1 {
		return s
	}
	for i := 0; i < l; i++ {
		j := arm1
		if i-j >= 0 && i+j < l && arr[i-j] == arr[i+j] {
			// 如果目前两边最长的范围都是ok的，再考虑会更长
			for j := 0; i-j >= 0 && i+j < l && arr[i-j] == arr[i+j]; j++ {
				if j > arm1 {
					pos1 = i
					arm1 = j
				}
			}
		}
		k := arm2
		if i > 0 && i+k-1 < l && i-k >= 0 && arr[i+k-1] == arr[i-k] {
			for k := 0; i+k-1 < l && i-k >= 0 && arr[i+k-1] == arr[i-k]; k++ {
				if k > arm2 {
					pos2 = i
					arm2 = k
				}
			}
		}
	}
	l1 := arm1*2 + 1
	l2 := arm2 * 2
	if l1 > l2 {
		return string(arr[pos1-arm1 : pos1+arm1+1])
	} else {
		return string(arr[pos2-arm2 : pos2+arm2])
	}
}
