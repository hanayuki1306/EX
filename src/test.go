package main

import "fmt"

func beautifulArray(N int) []int {
	var A []int
	A = []int{1}
	for len(A) < N {
		for i := 0; i < len(A); i++ {
			if i*2-1 <= N {
				A = append(A, (i*2 - 1))
			}
			for i := 0; i < len(A); i++ {
				if i*2 <= N {
					A = append(A, (i * 2))
				}
			}
		}
		break
	}
	return A
}
func main() {
	a := 5
	fmt.Println(beautifulArray(a))
}
