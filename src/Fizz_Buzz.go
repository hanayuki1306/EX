// https://leetcode.com/problems/fizz-buzz/
package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			num := strconv.Itoa(i)
			res = append(res, num)
		}
	}
	return res
}
func main() {
	n := 14
	fmt.Println(fizzBuzz(n))
}
