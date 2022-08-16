package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [...]int{14, 1, 1, 1}
	n2 := nums[:]
	res := myMinSubArrayLen(7, n2)
	fmt.Println(res)
}
func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

func myMinSubArrayLen(target int, nums []int) int {
	l := len(nums)
	left, right, sum := 0, 0, 0
	ans := math.MaxInt
	for right < l-1 {
		sum += nums[right]
		for sum >= target {
			ans = Min(right-left+1, ans)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt {
		return 0
	} else {
		return ans
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
