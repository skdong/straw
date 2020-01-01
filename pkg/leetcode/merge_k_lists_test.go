package leetcode

import (
	"testing"
)

func intsLists(lists []*ListNode) [][]int {
	var vals [][]int
	vals = make([][]int, len(lists))
	for i := 0; i < len(lists); i++ {
		node := lists[i]
		for node != nil {
			vals[i] = append(vals[i], node.Val)
			node = node.Next
		}
	}
	return vals
}
func intsList(list *ListNode) []int {
	var val []int
	for list != nil {
		val = append(val, list.Val)
		list = list.Next
	}
	return val
}

func EqualList(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
func TestMergeKLists(t *testing.T) {

	var vals [][]int = [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	var except []int = []int{1, 1, 2, 3, 4, 4, 5, 6, 7}
	var lists []*ListNode
	lists = make([]*ListNode, len(vals))
	for i := 0; i < len(vals); i++ {
		head := &ListNode{Val: 0, Next: nil}
		node := head
		for j := 0; j < len(vals[i]); j++ {
			node.Next = &ListNode{Val: vals[i][j], Next: nil}
			node = node.Next
		}
		lists[i] = head.Next
	}
	list := mergeKLists(lists)

	if !EqualList(intsList(list), except) {
		t.Errorf("except:%v real: %v", except, intsList(list))
	}
}
