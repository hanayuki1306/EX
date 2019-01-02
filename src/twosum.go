package main 
import "fmt"

func twoSum(nums  []int, target2 int) [][]int {
	var result [][] int
	// var index: = 0
	for i := 0; i < len(nums); i++ {
		for j := i+1; j<len(nums); j++ {
			if nums[i] + nums[j] == target2 {
				// result = [2]int{i, j}
				result = append(result, []int{i,j})
				// fmt.Println(nums[i], nums[j])
			}
		}
	}
	return result
}

func main(){
	var nums [] int
	nums  = []int{2,5,6,9,10}
	target := 11
	fmt.Println(twoSum(nums, target))
}


//  https://leetcode.com/problems/two-sum/ 
