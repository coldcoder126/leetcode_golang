package main

import "fmt"

func main() {
	t := []int{1, 3, -1, -3, 5, 3, 6, 7}
	r := maxSlidingWindow2(t, 3)
	fmt.Println(r)
}

// 执行用时： 864 ms , 在所有 Go 提交中击败了 5.16% 的用户
// 内存消耗： 9.1 MB , 在所有 Go 提交中击败了 43.94% 的用户
// 通过测试用例： 51 / 51
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	left, right := 0, k-1
	curMax := nums[0]
	for i := 0; i < k; i++ {
		curMax = max(curMax, nums[i])
	}
	res = append(res, curMax)

	//如果暴力遍历，要k*n
	//在此基础上优化一下：
	//	1. 如果窗口left>=curMax，说明最大值可能要换了
	//	2. 如果窗口right>curMax，curMax = nums[right]，
	for right < len(nums)-1 {
		left++
		right++
		if nums[left-1] >= curMax && nums[left] < curMax && nums[right] < curMax {
			//需要在窗口中重新找最大值
			curMax = nums[left]
			for i := left; i <= right; i++ {
				curMax = max(curMax, nums[i])
			}
			//continue
		} else if nums[right] >= curMax {
			curMax = nums[right]
			//continue
		}
		//否则最大值没变
		res = append(res, curMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/**
方式二：维护一个单调队列
1. pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
2. push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
*/

type MyQue struct {
	val []int
}

// 队列是否为空
func (q *MyQue) empty() bool {
	return len(q.val) == 0
}

func (q *MyQue) backVal() int {
	return q.val[len(q.val)-1]
}

func (q *MyQue) frontVal() int {
	return q.val[0]
}

// 从左侧弹出
func (q *MyQue) pop() int {
	cur := q.val[0]
	q.val = q.val[1:]
	return cur
}

// 从右侧进入
func (q *MyQue) push(x int) {
	for !q.empty() && x > q.backVal() {
		q.val = q.val[:len(q.val)-1]
	}
	q.val = append(q.val, x)
}

/*
*
执行用时：244 ms, 在所有 Go 提交中击败了 19.25%的用户
内存消耗：8.8 MB, 在所有 Go 提交中击败了83.31%的用户
通过测试用例：51 / 51
*/
func maxSlidingWindow2(nums []int, k int) []int {
	res := make([]int, 0)
	que := &MyQue{}
	left, right := 0, k-1
	for i := 0; i < k; i++ {
		que.push(nums[i])
	}
	res = append(res, que.val[0])

	for right < len(nums)-1 {
		left++
		right++
		if nums[left-1] == que.frontVal() {
			que.pop()
		}
		que.push(nums[right])
		res = append(res, que.frontVal())
	}
	return res
}
