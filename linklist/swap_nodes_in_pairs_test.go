package linklist

// https://leetcode.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head.Next

	var prevPrevNode *ListNode = nil
	prevNode := head
	curr := head.Next

	for curr != nil {
		if prevNode != nil {
			currNext := curr.Next
			prevNode.Next = currNext
			curr.Next = prevNode
			if prevPrevNode != nil {
				prevPrevNode.Next = curr
			}
			prevPrevNode = prevNode
			prevNode = prevNode.Next
			if prevNode == nil {
				curr = nil
			} else {
				curr = prevNode.Next
			}
		}
	}

	if newHead == nil {
		return head
	}

	return newHead
}
