package golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
	for _, test := range tests {
		t.Run("running-sum "+fmt.Sprint(test.input), func(t *testing.T) {
			actual := runningSum(test.input)
			assert.Equal(t, test.expect, actual)
		})
	}
}
