package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0, 1)
	s := &nodeStack{arr: make([]*TreeNode, 0, 1)}
	curr := root

	for curr != nil || s.len() != 0 {
		for curr != nil {
			s.push(curr)
			curr = curr.Left
		}

		curr = s.pop()
		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}

type nodeStack struct {
	arr []*TreeNode
}

func (s *nodeStack) len() int {
	return len(s.arr)
}

func (s *nodeStack) push(val *TreeNode) {
	s.arr = append(s.arr, val)
}

func (s *nodeStack) top() *TreeNode {
	return s.arr[len(s.arr)-1]
}

func (s *nodeStack) bottom() *TreeNode {
	return s.arr[0]
}

func (s *nodeStack) pop() *TreeNode {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}
