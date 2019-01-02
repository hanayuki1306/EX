package main 
import "fmt"

func main() {
	var A[5] int 
	A = [5]int{2,5,6,9,10}
	


	delete(A, 2 )
	fmt.Println(A)
	
}