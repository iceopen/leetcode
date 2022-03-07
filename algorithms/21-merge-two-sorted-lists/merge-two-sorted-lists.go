package merge_two_sorted_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 有一条链为nil，直接返回另一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	// 循环比较l1和l2的值，始终选择较小的值连上node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		// 有了这一步，head才是一个完整的链
		node = node.Next
	}

	if l1 != nil {
		// 连上l1剩余的链
		node.Next = l1
	}

	if l2 != nil {
		// 连上l2剩余的链
		node.Next = l2
	}

	return head
}

// 速度比较快的代码算法
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	ret := ListNode{}
// 	c := &ret

// 	for l1 != nil && l2 != nil {
// 		if l1.Val <= l2.Val {
// 			c.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			c.Next = l2
// 			l2 = l2.Next
// 		}
// 		c = c.Next
// 	}

// 	if l1 == nil {
// 		c.Next = l2
// 	} else if l2 == nil {
// 		c.Next = l1
// 	}

// 	return ret.Next
// }
