package main 
import "fmt"

func main() {
	var A[5] int 
	A = [5]int{2,5,6,9,10}
	

	for i := 0; i<len(A); i++ {
		for j := i+1;j<len(A);j++ {
			fmt.Println(A[i]+A[j])
		}
	}
	
}