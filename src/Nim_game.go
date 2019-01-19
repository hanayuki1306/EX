// https://leetcode.com/problems/nim-game/submissions/

package main

import "fmt"

// chi can ko chia het cho 4 la van co the win
func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}
	return true
}
func main() {
	fmt.Println(canWinNim(40))
}
