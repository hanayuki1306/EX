// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/submissions/
package main

import (
	"fmt"
	"sort"
)

// func minMoves(nums []int) int {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {

// 		for j := 0; j < n-1-i; j++ {
// 			if nums[j] > nums[j+1] {
// 				tmp :	= nums[j]
// 				nums[j] = nums[j+1]
// 				nums[j+1] = tmp
// 			}
// 		}
// 	}
// 	ans := nums[n-1] - nums[0]+ 1
// 	return ans
// }

func minMoves(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)

	count := 0
	lenNum := len(nums) - 1
	fmt.Println(len(nums))

	for i := 0; i < lenNum; i++ {
		count += (nums[i+1] - nums[i]) * (lenNum - i)

	}

	return count
}
func main() {
	a := []int{1, 3, 5, 4}
	fmt.Println(minMoves(a))
}
