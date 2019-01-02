package main

import (
	"fmt"
	"strconv"
)

func countAndsay(s string) (r string) {
	c := s[0]
	nc := 1
	for i := 1; i < len(s); i++ {
		d := s[i]
		if d == c {
			nc++
			fmt.Println(nc)
			continue
		}
		r += strconv.Itoa(nc) + string(c)
		c = d
		nc = 1
	}
	return r + strconv.Itoa(nc) + string(c)
}

func main() {
	s := "1"
	fmt.Println(s)
	for i := 0; i < 8; i++ {
		s = countAndsay(s)
		fmt.Println(s)
	}
}
