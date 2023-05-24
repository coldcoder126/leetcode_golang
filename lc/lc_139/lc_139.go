package main

import (
	"fmt"
	"math"
)

func main() {
	s := "bb"
	arr := []string{"a", "b", "bbb", "bbbb"}

	res := wordBreak(s, arr)
	fmt.Println(res)
}

func wordBreak(s string, wordDict []string) bool {
	//0-1背包模型，单词是物品，字符串是背包
	//dp[i]是到索引为i的字符串是否可以被合成
	l := len(s)
	dp := make([]bool, l+1)

	//组装成map
	mp := make(map[string]struct{})
	maxWordLen := 0
	minWorlLen := math.MaxInt32
	for i := 0; i < len(wordDict); i++ {
		mp[wordDict[i]] = struct{}{}
		maxWordLen = max(maxWordLen, len(wordDict[i]))
		minWorlLen = min(minWorlLen, len(wordDict[i]))
	}

	for k, _ := range mp {
		if len(k) <= l && k == s[0:len(k)] {
			dp[len(k)-1] = true
		}
	}

	for i := 0; i < l; i++ {
		for j := minWorlLen; j <= maxWordLen; j++ {
			if dp[i] && i+j < l {
				temp := s[i+1 : i+j+1]
				if _, ok := mp[temp]; ok {
					dp[i+j] = true
				}
			}
		}
	}

	return dp[l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
