package linklist

// https://leetcode.com/problems/linked-list-cycle-ii/
// https://cs.stackexchange.com/questions/10360/floyds-cycle-detection-algorithm-determining-the-starting-point-of-cycle
// https://medium.com/@ChYuan/leetcode-no-142-linked-list-cycle-ii-%E5%BF%83%E5%BE%97-medium-c5b53dcca804
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	mark := false

	for slow != nil && fast != nil && fast.Next != nil && !mark {
		slow = slow.Next
		fast = fast.Next.Next
		if slow != nil && slow == fast {
			mark = true
		}
	}

	if !mark {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
