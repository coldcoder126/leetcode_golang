package main

import "fmt"

func main() {
	a := letterCombinations("234")
	fmt.Print(a)
}

var mp = map[int][]byte{2: []byte{'a', 'b', 'c'},
	3: []byte{'d', 'e', 'f'},
	4: []byte{'g', 'h', 'i'},
	5: []byte{'j', 'k', 'l'},
	6: []byte{'m', 'n', 'o'},
	7: []byte{'p', 'q', 'r', 's'},
	8: []byte{'t', 'u', 'v'},
	9: []byte{'w', 'x', 'y', 'z'}}

func letterCombinations(digits string) []string {
	var res []byte
	var ans []string
	if len(digits) == 0 {
		return ans
	}

	var input []int
	for _, i := range digits {
		input = append(input, int(i-'0'))
	}
	dfs(input, 0, &res, &ans)
	return ans
}

func dfs(input []int, begin int, res *[]byte, ans *[]string) {

	if len(*res) == len(input) {
		temp := string(*res)
		*ans = append(*ans, temp)
		return
	}

	for i := begin; i < len(input); i++ {
		arr := mp[input[i]]
		for _, a := range arr {
			*res = append(*res, a)
			dfs(input, i+1, res, ans)
			*res = (*res)[:len(*res)-1]
		}
	}
}
