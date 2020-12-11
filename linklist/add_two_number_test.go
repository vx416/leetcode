package linklist

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}

	curr := root
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		ext := curr.Val
		curr.Val = (val1 + val2 + ext) % 10
		ext = (val1 + val2 + ext) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		// has next number
		if l1 != nil || l2 != nil {
			curr.Next = &ListNode{Val: ext}
			curr = curr.Next
		} else if ext != 0 {
			curr.Next = &ListNode{Val: ext}
		}
	}

	return root
}
