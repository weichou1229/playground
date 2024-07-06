package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name       string
		inputList1 *ListNode
		inputList2 *ListNode
		expected   *ListNode
	}{
		{
			"merge two list",
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			"merge empty lists", &ListNode{Val: 0, Next: nil}, nil, &ListNode{Val: 0, Next: nil},
		},
		{
			"one is empty list", nil, nil, nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := mergeTwoLists(test.inputList1, test.inputList2)
			assert.Equal(t, actual, test.expected)
		})
	}
}
