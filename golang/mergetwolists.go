package golang

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// !! use recursive can improve the performance

	var list []*ListNode

	for i := list1; ; i = i.Next {
		list = append(list, i)
		if i.Next == nil {
			break
		}
	}
	for i := list2; ; i = i.Next {
		list = append(list, i)
		if i.Next == nil {
			break
		}
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Val < list[j].Val
	})

	nextNode := list[0]
	length := len(list)
	for i := 1; i < length; i++ {
		nextNode.Next = list[i]
		nextNode = list[i]
	}
	return list[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
