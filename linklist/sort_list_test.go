package linklist

// https://leetcode.com/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next != nil && head.Next.Next == nil {
		if head.Val < head.Next.Val {
			return head
		}
		next := head.Next
		next.Next = head
		head.Next = nil
		return next
	}

	midNode := findMidNode(head)
	rightHead := midNode.Next
	midNode.Next = nil
	newLeftHead := sortList(head)
	newRightHead := sortList(rightHead)
	return merge(newLeftHead, newRightHead)
}

func findMidNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var (
		newHead *ListNode
		curr    *ListNode
	)

	l1Node, l1Head := pop(l1)
	l2Node, l2Head := pop(l2)

	for l1Node != nil && l2Node != nil {
		var appendNode *ListNode
		if l1Node.Val < l2Node.Val {
			appendNode = l1Node
			l1Node, l1Head = pop(l1Head)
		} else {
			appendNode = l2Node
			l2Node, l2Head = pop(l2Head)
		}

		if curr == nil {
			curr = appendNode
		} else {
			curr.Next = appendNode
			curr = curr.Next
		}

		if newHead == nil {
			newHead = curr
		}
	}

	if l1Node != nil {
		l1Node.Next = l1Head
		curr.Next = l1Node
	}

	if l2Node != nil {
		l2Node.Next = l2Head
		curr.Next = l2Node
	}

	return newHead
}

func pop(l *ListNode) (*ListNode, *ListNode) {
	if l == nil {
		return nil, nil
	}
	newHead := l.Next
	l.Next = nil

	return l, newHead
}
