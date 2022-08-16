package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {
		fmt.Printf("%v： %v \n", i, isHappy(i))
	}

}
func isHappy(n int) bool {
	slow, fast := n, getSum(n)
	for slow != fast {
		slow = getSum(slow)
		fast = getSum(fast)
		fast = getSum(fast)
	}
	return slow == 1

}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		bit := n % 10
		sum += bit * bit
		n = n / 10
	}
	return sum
}
