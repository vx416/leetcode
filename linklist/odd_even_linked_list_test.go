package linklist

// https://leetcode.com/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var (
		preEvenHead *ListNode = head
		evenHead    *ListNode = head.Next
		curr                  = head
		currPos               = 1
	)

	for curr != nil {
		if currPos%2 == 0 {
			evenCurr := curr
			oldCurr := evenCurr.Next
			if oldCurr != nil {
				oldCurrNext := oldCurr.Next
				oldCurr.Next = evenHead
				evenCurr.Next = oldCurrNext
				preEvenHead.Next = oldCurr
				preEvenHead = oldCurr
			} else {
				curr = nil
			}
		} else {
			curr = curr.Next
		}
		currPos++
	}

	return head
}
