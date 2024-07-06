package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		expected     int
		expectedList []int
	}{
		{"return 2", []int{1, 1, 2}, 2, []int{1, 2}},
		{"return 5", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeDuplicates(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
