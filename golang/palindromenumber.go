package golang

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
