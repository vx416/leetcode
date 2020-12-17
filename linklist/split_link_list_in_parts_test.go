package linklist

// https://leetcode.com/problems/split-linked-list-in-parts/
func splitListToParts(root *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	if root == nil {
		return res
	}

	n := listLen(root)

	m := n / k
	d := n % k
	curr := root
	for i := range res {
		head := curr
		l := m
		if i < d {
			l++
		}
		for j := 0; j < l-1; j++ {
			if curr == nil {
				break
			}
			curr = curr.Next
		}
		if curr != nil {
			tail := curr
			curr = tail.Next
			tail.Next = nil
		}
		res[i] = head
	}

	return res
}

func listLen(node *ListNode) int {
	n := 0

	curr := node
	for curr != nil {
		curr = curr.Next
		n++
	}

	return n

}
