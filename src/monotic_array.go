https://leetcode.com/problems/monotonic-array/

func isMonotonic(A []int) bool {
    if len(A) == 1 {
        return true
    }
    increase := true
    decrease := true
    for i := 0; i < len(A) - 1; i++ {        //i++ A[i] >A[i+1]
        if A[i] > A[i + 1] {
            decrease = false
        } else if A[i] < A[i + 1] {
            increase = false
        }
        if (increase || decrease) == false {
            return false
        } 
    }
    return true
}


// package main 
// import "fmt"
// func isMonotonic(A []int) bool {
//     if len(A) == 1 {
//         return true
//     }
//     increase := true
//     decrease := true
//     for i := 0; i < len(A) - 1; i++ {        //i++ A[i] >A[i+1] thi check increase (1)
//         if A[i] > A[i + 1] {
//             decrease = false
//         } else if A[i] < A[i + 1] {          // nugoc lai vs (1)
//             increase = false
//         }
//         if (increase || decrease) == false {
//             return false
//         } 
//     }
//     return true
// }
// func main() {
// 	A:= []int {1,2,3,10,6,7,8,10}
// 	fmt.Println(isMonotonic(A))
// }