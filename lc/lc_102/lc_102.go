package main

import (
	"container/list"
	"fmt"
	"leetcode/lc/util"
)

type TreeNode = util.TreeNode

func main() {
	arr := make([][]int, 0)
	arr1 := []int{1}
	arr2 := []int{4, 5}
	arr = append(arr, arr1)
	arr = append(arr, arr2)
	//copy(arr1, arr2)
	fmt.Println(arr)

}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		// 1. 将que中的所有取出
		temp := make([]*TreeNode, 0)
		vals := make([]int, 0)
		for i := 0; i < len(que); i++ {
			node := que[i]
			vals = append(vals, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		res = append(res, vals)
		que = temp
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}
	return res
}
