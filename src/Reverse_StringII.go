package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// func reverseStr(s string, k int) string {
// 	a := len(s)
// 	if a > k {
// 		Reverse(s)
// 	}
// 	return Reverse(s)
// }
func main() {
	s := "motcauhoikhokhan"
	k := 3
	fmt.Println(reverseString(s, k))

}

// for 1 vong tu 1 den k
