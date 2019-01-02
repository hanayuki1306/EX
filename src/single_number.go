 // a.31
 package main 
 import "fmt"
 func singleNumber(nums []int) int {
	var key int
	for _, num:= range nums {
		key ^= num
	}
	return key
 }

 func main() {
	 var dauvao []int
	 dauvao = []int {5,3,2,1,4,4,3,5,2}
// fmt.Println(singleNumber())
 fmt.Println(singleNumber(dauvao))
 }

//  https://leetcode.com/problems/single-number/submissions/
  //toan hang XOR 
//  https://vietjack.com/cplusplus/toan_tu_trong_cplusplus.jsp  ^=