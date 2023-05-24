package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "101023"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了69.37%的用户
*/

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	ans := make([]string, 0)
	var dfs func(substr string)
	dfs = func(substr string) {
		//如果substr符合要求，并且len(ans) == 3，则符合结束条件
		if check(substr) && len(ans) == 3 {
			if x, err := strconv.Atoi(substr); err == nil && x <= 255 {
				ans = append(ans, substr)
				temp := strings.Join(ans, ".")
				res = append(res, temp)
				ans = ans[:3]
			}
		}
		sublen := len(substr)
		for i := 1; i < sublen; i++ {
			s1 := substr[:i]
			if check(s1) && len(ans) < 4 {
				if x, err := strconv.Atoi(s1); err == nil && x <= 255 {
					ans = append(ans, s1)
					s2 := substr[i:]
					dfs(s2)
					ans = ans[:len(ans)-1]
				}
			}
		}
	}
	dfs(s)
	return res
}

func check(s string) bool {
	if len(s) == 1 || (len(s) > 1 && len(s) <= 3 && !strings.HasPrefix(s, "0")) {
		return true
	}
	return false
}
