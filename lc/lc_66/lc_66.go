package main

func plusOne(digits []int) []int {

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		cur := digits[i] + carry
		if carry == 1 {
			if cur == 10 {
				carry = 1
				cur = 0
			} else {
				carry = 0
			}
			digits[i] = cur
		} else {
			break
		}
	}
	if carry == 1 {
		arr := []int{1}
		arr = append(arr, digits...)
		return arr
	}
	return digits
}
