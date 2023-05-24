package main

import "fmt"

func main() {
	s := "abbc"
	res := partition(s)
	fmt.Println(res)
}

/**
执行用时：220 ms, 在所有 Go 提交中击败了94.62%的用户
内存消耗：23 MB, 在所有 Go 提交中击败了55.15%的用户
*/

// 回溯和递归
func partition(s string) [][]string {
	res := make([][]string, 0)
	ans := make([]string, 0)
	var dfs func(substr string)
	dfs = func(substr string) {
		if check(substr) {
			ans = append(ans, substr)
			temp := make([]string, len(ans))
			copy(temp, ans)
			res = append(res, temp)
			ans = ans[:len(ans)-1]
			if len(substr) == 1 {
				return
			}
		}

		for i := 1; i < len(substr); i++ {
			s1 := substr[:i]
			if check(s1) {
				ans = append(ans, s1)
				s2 := substr[i:]
				dfs(s2)
				ans = ans[:len(ans)-1]
			} else {
				continue
			}
		}
	}
	dfs(s)
	return res
}

// 检查一个字符串是否是回文子串
func check(s string) bool {
	b := []byte(s)
	l := len(b)
	for i := 0; i < l/2; i++ {
		if b[i] == b[l-1-i] {
			continue
		} else {
			return false
		}
	}
	return true
}
