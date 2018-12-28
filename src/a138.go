package main 
import "fmt"
import "strconv"

func countAndSay(n int) string {

	s:= strconv.Itoa(n)	
	return s
}
// func nhap_vao(i string ) int {
// 	i := strconv.Atoi("1") return 
// }
func main() {
	// s:= strconv.Itoa(12)
	// fmt.Println(s)
	// fmt.Printf("%T",s)
	// var a,b,c,d,e int
	// a,b,c,d,e =  1,11,21,1211,111221
	a:=1211
	fmt.Println(countAndSay(a))
}
