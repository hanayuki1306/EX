// <97 chu hoa
// >97 chu thuong quy theo byte
// https://leetcode.com/problems/detect-capital/submissions/
package main

import (
	"fmt"
	"unicode"
)

func detectCapitalUse(word string) bool {
	lower := 0
	upper := 0
	for _, letter := range word {
		if unicode.IsUpper(letter) && lower > 0 {
			return false
		} else if unicode.IsLower(letter) && upper > 1 {
			return false
		} else if unicode.IsLower(letter) {
			lower++
		} else if unicode.IsUpper(letter) {
			upper++
		}
	}
	return true
}
func main() {
	a := "flagG"
	fmt.Println(detectCapitalUse(a))
}

// func detectCapitalUse(word string) bool {
//     if isCapital(word[0]){
//         for i:=1 ; i < len(word) - 1; i++{
//             if isCapital(word[i]) != isCapital(word[i+1]){
//                 return false
//             }
//         }
//     }else{
//         for i:=1 ; i < len(word); i++{
//             if isCapital(word[i]){
//                 return false
//             }
//         }
//     }
//     return true
// }

// func isCapital(b byte) bool{
//     return b >= 65 && b <= 90
// }

// package main

// import "fmt"

// func detectCapitalUse(word string) bool {
// 	ans := []byte(word)

// 	lower := 0
// 	upper := 0
// 	for _, v := range ans {
// 		if ans[v] < 97 {
// 			upper++
// 		} else {
// 			lower++
// 		}
// 	}
// 	for _, v := range ans {
// 		if ans[v	] < 97 && lower > 0 {
// 			return false
// 		} else if ans[v] > 97 && upper > 1 {
// 			return false
// 		} else if ans[v] > 97 {
// 			lower++
// 		} else if ans[v] < 97 {
// 			upper++
// 		}
// 	}
// 	return true
// }
// func main() {
// 	a := "USA"
// 	fmt.Println(detectCapitalUse(a))
// }
