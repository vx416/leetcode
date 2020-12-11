package graph

import (
	"testing"
)

func TestZigZag(t *testing.T) {
	t.Log(zigzagLevelOrder(arrToTree([]int{1, 2, 3, 4, -1, -1, 5})))
	t.Log(zigzagLevelOrder(arrToTree([]int{3, 9, 20, -1, -1, 15, 7})))
}

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
	stack := &nodeStack{arr: make([]*TreeNode, 0, 1)}
	stack.push(root)
	res := make([][]int, 0, 1)
	zigZagbst(stack, &res, 1)

	return res
}

func zigZagbst(stack *nodeStack, res *[][]int, level int) {
	if stack.isEmpty() {
		return
	}
	subRes := make([]int, 0, 1)
	newStack := &nodeStack{arr: make([]*TreeNode, 0, 1)}
	for !stack.isEmpty() {
		subN := stack.pop()
		subRes = append(subRes, subN.Val)
		if level%2 == 1 {
			newStack.push(subN.Left)
			newStack.push(subN.Right)
		} else {
			newStack.push(subN.Right)
			newStack.push(subN.Left)
		}
	}
	*res = append(*res, subRes)
	zigZagbst(newStack, res, level+1)
}

type nodeStack struct {
	arr []*TreeNode
}

func (s *nodeStack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *nodeStack) push(n *TreeNode) {
	if n == nil {
		return
	}
	s.arr = append(s.arr, n)
}

func (s *nodeStack) pop() *TreeNode {
	n := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return n
}
