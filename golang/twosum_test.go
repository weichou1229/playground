package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	nums   []int
	target int
	want   []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		t.Run("two-sum "+fmt.Sprint(test.nums), func(t *testing.T) {
			out := twoSum(test.nums, test.target)
			assert.ElementsMatch(t, test.want, out)
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for _, test := range tests {
		for i := 0; i < b.N; i++ {
			twoSum(test.nums, test.target)
		}
	}

}
