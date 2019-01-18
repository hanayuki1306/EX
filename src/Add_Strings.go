// https://leetcode.com/problems/add-strings/submissions/

// package main

// import "fmt"

// func main() {
// 	s := "99888"
// 	a := []byte(s)
// 	fmt.Println(a)
// 	a1 := 99 % 10
// 	fmt.Println(a1)
// }

package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	st1, st2 := []byte(num1), []byte(num2)

	if len(st1) < len(st2) {
		st1, st2 = st2, st1
	}
	sum := byte(0)
	for i, j := len(st1)-1, len(st2)-1; i >= 0 || j >= 0; i, sum = i-1, sum/10 { // for i>=0 j:=len s1||for j>=0 j= len s2-1
		if j >= 0 {
			sum += st2[j] - '0'
			j--
		}
		sum += st1[i] - '0'
		st1[i] = sum%10 + '0'
	}

	if (sum) != byte(0) {
		st1 = append([]byte{'1'}, st1...)
	}
	return string(st1)
}
func main() {
	a := "987"
	b := "0"
	fmt.Println(addStrings(a, b))
}

// package main
// import "fmt"
// func addString(a string, b string) string {
// 	s1, s2:= []byte(a) , []byte(b)
// 	sum := byte(0)

// 	for i;j := len(s1)-1, len(s2)-1; i>=0  || j>=0; i=i-1, sum = sum/10 {
// 		if j>0 {
// 			sum += s2[j] - '0'
// 			j--
// 		}
// 	}

// }
