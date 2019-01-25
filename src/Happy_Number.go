// https://leetcode.com/problems/happy-number/
package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n != 1 {
		n = sub(n)
		if n == 4 {
			return false
		}
	}
	return true
}

func sub(n int) int {
	str := fmt.Sprintf("%d", n)
	fmt.Println(str)
	var ret int
	for _, v := range str {
		ret += int(math.Pow(float64(v-'0'), 2))

	}
	return ret
}
func main() {
	fmt.Println(isHappy(82))
}
