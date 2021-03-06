// https://leetcode.com/problems/maximum-average-subarray-i/submissions/
package main

import "fmt"

// func findMaxAverage(nums []int, k int) float64 {

// 	sum := make([]int, len(nums)+1)
// 	for i, v := range nums {
// 		sum[i+1] = sum[i] + v
// 	}
// 	max := -1000000000000000
// 	for i := 0; i+k < len(sum); i++ {
// 		if sum[i+k]-sum[i] > max {
// 			max = sum[i+k] - sum[i]
// 		}
// 	}
// 	return float64(max) / float64(k)
// }

func findMaxAverage(nums []int, k int) float64 {
	result := 0
	// @TODO: result will become ...
	for i := 0; i < k; i++ {
		result += nums[i]
	}

	for i := k; i < len(nums); i++ {
		resultTmp := 0

		for j := i; j > i-k; j-- {
			resultTmp += nums[j]
		}
		if resultTmp > result {
			result = resultTmp
		}
	}

	return float64(result) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}

// 24
// i=4 <8
// anst:=5
// for j=5;j>1
