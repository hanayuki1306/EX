package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var key byte
	for i := 0; i < len(s); i++ {
		key ^= s[i]
		// key = s[i]+ key1
		fmt.Println("key loop1 = ", s[i])
	}

	for j := 0; j < len(t); j++ {
		key ^= t[j]
	}
	// fmt.Println("key loop2 =", key)
	return key
}
func main() {
	sida := "string"
	thangsida := "strings"
	fmt.Printf("%T\n\n\n", findTheDifference(sida, thangsida))
	s := findTheDifference(sida, thangsida)
	fmt.Println(string(s))
}
