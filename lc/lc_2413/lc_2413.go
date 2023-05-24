package main

import "fmt"

func main() {
	res := smallestEvenMultiple(5)
	fmt.Println(res)
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return n * 2
	}
}
