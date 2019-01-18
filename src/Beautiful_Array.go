package main

import "fmt"

func beautifulArray(N int) []int {
	temp := 0
	var A []int
	for i := 1; i < N+1; i++ {
		A = append(A, i)
	}
	fmt.Println(A)
	for i := 0; i < len(A); i++ {
		var k, j int
		count := 2
		C := 1
		k = i + C

		for j = (i + count); j < len(A); j++ {
			if (A[k] * 2) == (A[i] + A[j]) {
				temp = A[i]
				A[i] = A[k]
				A[k] = temp
				fmt.Printf("count=%d\n", count)

			}
			C += 1
			if (A[k] * 2) == (A[i] + A[j]) {
				temp = A[i]
				A[i] = A[k]
				A[k] = temp
			}

			count += 1

			fmt.Printf("C=%d\n\n", C)
			fmt.Println(A)
		}

	}
	return A
}
func main() {
	fmt.Println(beautifulArray(5))
}
