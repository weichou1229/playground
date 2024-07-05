package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"1", "()", true},
		{"2", "()[]{}()", true},
		{"3", "(]", false},
		{"4", "(123])", false},
		{"5", "(([]test))", true},
		{"6", "([)]", false},
		{"7", "(){}}{", false},
		{"8", "({{{{}}}))", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isValidParentheses(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
