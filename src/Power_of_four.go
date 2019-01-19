
https://leetcode.com/problems/power-of-four/submissions/

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	} else if n%4 != 0 || n == 0 {
		return false

	} else {
		for n != 0 {
			n = n / 4
			if n != 1 && n%4 != 0 {
				return false
			}
		}
	}
	return true
}