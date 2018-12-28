package main 
import "fmt"
func numJewelsInStones(J string, S string) int {
	var counter = 0
	for i:=0; i<len(S); i++{
		for j:=0; j<len(J); j++{
			if S[i] == J[j] { fmt.Println(S[i])
				 counter++
				//  fmt.Println(counter)
			}
		}
	}
	return counter
} 
func main() {
	var J,S string 
	J = "aAC"
	S = "aAAbbbbCC"
   fmt.Println(numJewelsInStones(J,S))
}
