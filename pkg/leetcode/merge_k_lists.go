package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	amount := len(lists)
	var merge2Lists func(l1, l2 *ListNode) *ListNode
	merge2Lists = func(l1, l2 *ListNode) *ListNode {
		head := ListNode{0, nil}
		node := &head
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				node.Next = l1
				l1 = l1.Next
			} else {
				node.Next = l2
				l2 = l2.Next
			}
			node = node.Next
			node.Next = nil
		}
		if l1 != nil {
			node.Next = l1
		}
		if l2 != nil {
			node.Next = l2
		}
		return head.Next
	}
	for interval := 1; interval < amount; interval *= 2 {
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
	}
	if amount == 0 {
		var list *ListNode
		return list
	} else {
		return lists[0]
	}
}
