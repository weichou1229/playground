package golang

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// You are given the heads of two sorted linked lists list1 and list2, return the head of the merged linked list.
// https://leetcode.com/problems/merge-two-sorted-lists
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
