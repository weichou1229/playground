package golang

// isPalindrome given an integer x, return true if x is a palindrome, and false otherwise.
// Palindrome: An integer is a palindrome when it reads the same forward and backward.
// https://leetcode.com/problems/palindrome-number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var slc []int
	for x > 0 {
		slc = append(slc, x%10)
		x = x / 10
	}

	for i := 0; i < len(slc)/2; i++ {
		if slc[i] != slc[len(slc)-i-1] {
			return false
		}
	}
	return true
}
