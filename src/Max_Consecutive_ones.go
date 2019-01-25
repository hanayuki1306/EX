// https://leetcode.com/problems/max-consecutive-ones/submissions/

package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for _, v := range nums {
		if v == 1 {
			count = count + 1
		} else {
			if count > max {
				max = count
			}
			count = 0

		}

	}
	if count > max {
		max = count
	}
	return max
}

func main() {
	a := []int{1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(a))
}
