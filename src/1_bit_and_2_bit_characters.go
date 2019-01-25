// https://leetcode.com/problems/1-bit-and-2-bit-characters/submissions/

package main

import "fmt"

// @TODO: Change the name 'ans' to 'isOneBitCharacter'
func ans(bits []int) bool {
	if len(bits) == 0 || bits[len(bits)-1] == 1 {
		return false
	}
	if len(bits) == 1 && bits[0] == 0 {
		return true
	}
	count := 0
	l := len(bits)
	for i := l - 2; i >= 0; i-- {
		if bits[i] == 1 {
			count += 1
		} else {
			break
		}
	}
	if count%2 == 0 {
		return true
	}
	return false
}

func main() {
	// @TODO: Should change 'a' to 'sampleArray'
	a := []int{1, 0, 1, 0, 0, 0, 1, 0, 0}
	fmt.Println(ans(a))
}

/* l == 0 hoac bit[l-1] ==1 false  vi chac chan 2 bit
   l == 1 bit [o] = 0 true

*/
