package main

import "fmt"

func singleNumber(nums []int) int {
	var key int
	for _, num := range nums {
		key ^= num
	}
	return key
}
func main() {
	nums := []int{2, 2, 3, 3, 1}
	fmt.Println(singleNumber(nums))

}

//   https://leetcode.com/problems/single-number/
