// https://leetcode.com/problems/132-pattern/submissions/
// package main

// import "fmt"

// func find132pattern(nums []int) bool {
// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[k] > nums[i] && nums[j] > nums[k] {
// 					return true
// 				}
// 			}
// 		}

// 	}
// 	return false
// }
// func main() {
// 	nums := []int{1, 2, 3, 4, 55, 6}
// 	fmt.Println(find132pattern(nums))
// 	//true
// }
// Time Limit Exceeded

package main

import (
	"fmt"
)

func min(a []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] > a[i+1] {
			temp := a[i]
			a[i] = a[i+1]
			a[i+1] = temp

		}

		return a[0]
	// }
}
func find132pattern(nums []int) bool {
	var min_i int

	// Change j to another
	for j := 0; j < len(nums)-1; j++ {
		min_i = min(min_i, nums[j])
	}
	for k := j + 1; k < len(nums); k++ {
		if nums[k] < nums[j] && min_i < nums[k] {
			return true
		}
	}

	return false
}
func main() {
	nums := []int{1, 2, 3, 4, 55, 6}
	fmt.Println(find132pattern(nums))
	//true
}
