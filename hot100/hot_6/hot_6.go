package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	res := convert(s, 7)
	fmt.Println(res)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := []byte(s)
	l := len(s)
	split_rows := make([][]byte, 0)
	res := make([]string, 0)
	step := 2*numRows - 2
	//1. 将arr按照每段 2N-2分开
	for idx := 0; idx < l; idx = idx + step {
		if idx+step >= l {
			split_rows = append(split_rows, arr[idx:])
		} else {
			split_rows = append(split_rows, arr[idx:idx+step])
		}
	}

	//2. 取数组指定位置，组成行
	for i := 0; i < numRows; i++ {
		line := make([]byte, 0)
		for j := 0; j < len(split_rows); j++ {
			if i == 0 || i == numRows-1 {
				if len(split_rows[j]) > i {
					line = append(line, split_rows[j][i])
				}
			} else {
				if len(split_rows[j]) > i {
					line = append(line, split_rows[j][i])
				}
				if len(split_rows[j]) > 2*(numRows-1)-i {
					line = append(line, split_rows[j][2*(numRows-1)-i])
				}
			}
		}
		s := string(line)
		res = append(res, s)
	}
	res_str := ""
	for _, sub_str := range res {
		res_str += sub_str
	}
	return res_str
}
