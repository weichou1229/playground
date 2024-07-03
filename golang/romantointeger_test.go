package golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, test := range tests {
		t.Run("roman-to-integer "+fmt.Sprint(test.input), func(t *testing.T) {
			actual := romanToInt(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
