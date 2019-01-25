// https://leetcode.com/problems/plus-one/

package main

import "fmt"

func plusOne(digits []int) []int {
	var ans int

	for i := len(digits) - 1; i >= 0; i-- {
		ans = digits[i] + 1
		if ans < 10 {
			digits[i] = ans
			return digits
		}

		digits[i] = ans - 10
	}

	return append([]int{1}, digits...)
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(plusOne(a))
}
