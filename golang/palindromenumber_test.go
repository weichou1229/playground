package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var isPalindromeTests = []struct {
	input    int
	expected bool
}{
	{121, true},
	{-121, false},
	{10, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range isPalindromeTests {
		t.Run("is-palindrome "+fmt.Sprint(test.input), func(t *testing.T) {
			res := isPalindrome(test.input)
			assert.Equal(t, test.expected, res)
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for _, test := range isPalindromeTests {
		for i := 0; i < b.N; i++ {
			isPalindrome(test.input)
		}
	}
}
