	
https://leetcode.com/problems/palindrome-number/


func isPalindrome(x int) bool {
	B:= strconv.Itoa(x)
	n:=len(B)
	for i:=0;i<len(B);i++ {
		if B[i]==B[n-1-i] { 
			continue
			
		} else {
			return false
			
		}
	}
	return true
}  

// package main 
// import ("fmt"
// 		"strconv"
// 	   )

// func bullshit(x int) bool {
// 	B:= strconv.Itoa(x)
// 	n:=len(B)
// 	for i:=0;i<len(B);i++ {
// 		if B[i]==B[n-1-i] { 
// 			continue
			
// 		} else {
// 			return false
			
// 		}
// 	}
// 	return true
// }


// //}
// func main() {
// 	A:= 12213221
// 	B:=strconv.Itoa(A)
// 	fmt.Println(A)
// 	fmt.Println(B)
// 	fmt.Println(bullshit(A))
// //	fmt.Println(B[0])
// //fmt.Println(B[6])
// }
