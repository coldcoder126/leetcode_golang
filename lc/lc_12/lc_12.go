package main

import (
	"sort"
	"strings"
)

func main() {
	//res := intToRoman(1993)
	//fmt.Println(res)
}

func intToRoman(num int) string {
	mp := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	sortedKey := make([]int, 0)
	for k, _ := range mp {
		sortedKey = append(sortedKey, k)
	}
	sort.Slice(sortedKey, func(i, j int) bool {
		return sortedKey[i] > sortedKey[j]
	})

	ans := make([]string, 0)
	idx := 0
	for num > 0 {
		for i := idx; i < len(sortedKey); i++ {
			k := sortedKey[i]
			if num >= k {
				ans = append(ans, mp[k])
				num = num - k
				idx = i
				break
			}
		}
	}
	return strings.Join(ans, "")
}
