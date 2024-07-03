package golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"aa", "ab"}, "a"},
	}
	for _, test := range tests {
		t.Run("LongestCommonPrefix "+fmt.Sprint(test.input), func(t *testing.T) {
			actual := longestCommonPrefix(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
