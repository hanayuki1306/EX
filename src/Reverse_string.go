// https://leetcode.com/problems/reverse-string/submissions/

package main

import "fmt"

// func reverseString(s string) string {
// 	r := []rune(s)
// 	l := len(s)
// 	for i := 0; i < l/2; i++ {
// 		r[i], r[l-1-i] = r[l-1-i], r[i]
// 	}
// 	return string(r)
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}
func main() {
	a := "sida"
	fmt.Println(reverseString(a))
}

// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
