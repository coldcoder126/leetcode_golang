package main

// 位运算
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

// 二进制中，奇数一定比前面的偶数多一个1；偶数x一定和 x/2 中1的数量相同
func countBits2(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}
