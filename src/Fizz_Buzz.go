// https://leetcode.com/problems/fizz-buzz/
package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		switch {
		// @TODO Should tab in front of 'case' and have space 'i%15' => 'i % 15'
		case i%15 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			num := strconv.Itoa(i)
			result = append(result, num)
		}
	}
	return result
}
func main() {
	n := 14
	fmt.Println(fizzBuzz(n))
}
