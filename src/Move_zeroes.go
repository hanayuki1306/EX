// https://leetcode.com/problems/move-zeroes/submissions/

package main

import "fmt"

func moveZeroes(nums []int) []int {
	// @TODO
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
	}
	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}
