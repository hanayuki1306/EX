// https://leetcode.com/problems/longest-palindrome/submissions/

package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, u := range s {
		m[u]++
	}
	var ketqua int
	var odd bool
	for _, v := range m {
		ketqua += v
		if v%2 == 1 {
			ketqua--
			odd = true
		}
	}
	if odd {
		ketqua++
	}
	return ketqua
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
