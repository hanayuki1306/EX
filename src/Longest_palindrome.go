// https://leetcode.com/problems/longest-palindrome/submissions/

package main

import "fmt"

func longestPalindrome(s string) int {

	// @TODO: m should named meanful
	m := make(map[rune]int)
	for _, u := range s {
		m[u]++
	}
	var result int
	var odd bool
	for _, v := range m {
		result += v
		if v%2 == 1 {
			result--
			odd = true
		}
	}
	if odd {
		result++
	}
	return result
}
func main() {
	a := "aabbccdd"
	fmt.Println(longestPalindrome(a))
}

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // func count(str string, c int32) int {
// // 	count := 0
// // 	for i := 0; i < len(str); i++ {
// // 		if int32(str[i]) == c {
// // 			count++
// // 		}
// // 	}
// // 	return count
// // }
// func main() {
// 	str := "a long string with many repeated characters"
// 	numberOfa := strings.Count(str, "a")

// 	fmt.Printf("[%v] string has %d of characters of [a] \n", str, numberOfa)
// 	fmt.Println(numberOfa)
// 	// fmt.Println(count(str, 'a'))
// }
