// https://leetcode.com/problems/ugly-number/submissions/
package main

import "fmt"

func isUgly(num int) bool {
	if num != 0 && num%2 == 0 {
		for num%2 == 0 {
			num = num / 2
		}
	}
	if num != 0 && num%3 == 0 {
		for num%3 == 0 {
			num = num / 3
		}
	}
	if num != 0 && num%5 == 0 {
		for num%5 == 0 {
			num = num / 5
		}
	}
	if num == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isUgly(1))
}
