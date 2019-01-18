// https://leetcode.com/problems/self-dividing-numbers/submissions/

package main

import "fmt"

// lay so du  = 0 ; val !=0
func so_thoa_man(gia_tri int) bool {
	num := gia_tri
	for num > 0 {
		chu_so := num % 10
		if chu_so == 0 || gia_tri%chu_so != 0 {
			return false
		}
		num = num / 10
	}
	return true
}
func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for num := left; num <= right; num++ {
		if so_thoa_man(num) {
			ans = append(ans, num)
		}
	}
	return ans
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
