
// https://leetcode.com/problems/power-of-two/

// 1 case n=1 true
//1 case n/2 khac = false
//
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if n%2 != 0 || n == 0 {
		return false

	} else {
		for n != 0 {
			n = n / 2
			if n != 1 && n%2 != 0 {
				return false
			}
		}
	}
	return true
}




