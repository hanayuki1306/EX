
// https://leetcode.com/problems/add-digits/

func addDigits(num int) int {
	if num <= 9 {
		return num
	}

	newNum := 0
	for num > 9 {
		newNum += num % 10
		num /= 10
	}
	newNum = newNum + num
	return addDigits(newNum)
}