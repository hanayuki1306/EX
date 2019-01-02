package main 
import "fmt"

func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
 }

func nextGreaterElement(findNums []int, num []int) []int {
	mang:= make ([]int,5)
	// var mang []int
	for i := range mang {
		mang[i] = -1
	}
	for k, v := range findNums {
		index := indexOf(v, num)
		if index == len(num)-1 {
			continue
		} 
		for i:=index+1; i<len(num)-1;i++ {
			
			if num[i] > v {
				mang[k] = num[i]
				} 
		}
		
		
		}
		return mang
	}
	

func main() {
	 num1:= []int {1,2,5,2,4}
	 num2:= []int {6,5,4,3,2,1,7}

	 fmt.Println(nextGreaterElement(num1,num2))
}



//   https://leetcode.com/problems/next-greater-element-i/submissions/