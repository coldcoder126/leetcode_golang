package main

import "fmt"

func main() {
	res := findAnagrams("cbaebabacd", "abc")
	fmt.Println(res)
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return res
	}

	var sCount, pCount [26]int

	//统计数量
	for i, _ := range p {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}
	//go中，两个数组中所有元素都是相等的时候数组才是相等的
	if sCount == pCount {
		res = append(res, 0)
	}

	//滑动窗口一直保持窗口大小
	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			res = append(res, i+1)
		}
	}

	return res
}
