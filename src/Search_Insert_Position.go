// https://leetcode.com/problems/search-insert-position/

package main

import (
	"fmt"
	"sort"
)

func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right + 1
}
func main() {
	a := []int{1, 2, 3, 1, 2, 3, 5}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(searchInsert(a, 3))
}
