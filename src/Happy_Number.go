// https://leetcode.com/problems/happy-number/
package main

import (
	"fmt"
	"math"
)

func sub(n int) int {
	str := fmt.Sprintf("%d", n)
	fmt.Println(str)
	var ret int
	for _, v := range str {
		ret += int(math.Pow(float64(v-'0'), 2))

	}
	return ret
}

func isHappy(number int) bool {
	if number <= 0 {
		return false
	}
	if number == 1 {
		return true
	}
	for number != 1 {
		number = sub(number)
		if number == 4 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isHappy(82))
}
