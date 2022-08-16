package main

import (
	"fmt"
	"sort"
)

func main() {
	ts := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(ts)
	fmt.Println(ans)
}

func groupAnagrams(strs []string) [][]string {
	// 排序
	res := make(map[string][]string)
	for _, s := range strs {
		arr := []byte(s)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
		s2 := string(arr)
		res[s2] = append(res[s2], s2)
	}

	ans := make([][]string, 0, len(res))
	for _, val := range res {
		ans = append(ans, val)
	}
	return ans
}

func groupAnagrams2(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
