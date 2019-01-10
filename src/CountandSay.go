//   https://leetcode.com/problems/count-and-say/submissions/

package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if 1 == n {
		return "1"
	}

	str1 := countAndSay(n - 1)
	len1 := len(str1)

	var ret string
	var c byte

	c = str1[0]
	count := 1
	for i := 1; i < len1; i++ {
		if str1[i] == c {
			count++
			continue
		}
		// ret = ret + string(count) + string(c)
		ret = ret + strconv.Itoa(count) + string(c)
		c = str1[i]
		count = 1
	}
	// ret = ret + string(count) + string(c)
	ret = ret + strconv.Itoa(count) + string(c)
	return ret
}

func main() {

	// for i := 0; i < 8; i++ {
	// 	s := countAndSay(i)
	// 	fmt.Println(s)

	// }
	// fmt.Println(countAndSay(0))
	fmt.Println(strconv.Itoa(65))
	fmt.Println(string('A'))
	for j := 1; j < 8; j++ {
		fmt.Println(countAndSay(j))

	}
}

// package main

// import "fmt"

// func main() {
// 	var s string
// 	s = "1111"
// 	// c := string(s[0])
// 	var k int
// 	k = 12
// 	fmt.Printf("%T\n", string(k))
// 	fmt.Printf("%T\n", s[1])
// 	fmt.Printf("%T\n", string(s[1]))

// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func countAndsay(s string) (r string) {
// 	c := s[0]
// 	nc := 1
// 	for i := 1; i < len(s); i++ {
// 		d := s[i]
// 		if d == c {
// 			nc++
// 			// fmt.Println(nc)
// 			continue
// 		}
// 		// r += string(nc) + string(c)

// 		r += strconv.Itoa(nc) + string(c)
// 		c = d
// 		nc = 1
// 	}
// 	// fmt.Printf("%T", string(c))
// 	return r + strconv.Itoa(nc) + string(c)

// }

// func main() {
// 	s := "1"
// 	fmt.Println(s)
// 	for i := 0; i < 8; i++ {
// 		s = countAndsay(s)
// 		fmt.Println(s)
// 	}
// }

// 1
// 11
// 21
// 1211
// 111221
// 312211
// 13112221
// 1113213211
// 31131211131221
