
https://leetcode.com/problems/missing-number/


func missingNumber(nums []int) int {
    A := len(nums)
    var B int
    for i:=0 ; i< A;i++{
        B = B ^ i ^ nums[i]
    }
    return B^A
}



//package main 
//import "fmt"
//func missing(nums []int) int {
	//A:= len(nums)
	//var B int
	//for i:=0; i < A; i++ {
		//B = B ^ i ^ nums[i]
	//}
//	return B^A
//}
//func main() {
//	A:= []int{0,1,2,3,6,4,7}
//
//	fmt.Println(A)
//	
//	fmt.Println(missing(A))
//}