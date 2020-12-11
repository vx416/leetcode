package tree

import (
	"testing"
)

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// preorder -> VLR
// inorder -> LVR

func TestBuildTree(t *testing.T) {
	t.Log(buildTree([]int{1, 2}, []int{1, 2}))
	t.Log(buildTree([]int{1, 2, 3}, []int{1, 2, 3}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	root := preorder[0]
	rootIndex := 0
	for i := range inorder {
		if inorder[i] == root {
			rootIndex = i
			break
		}
	}
	rootNode := &TreeNode{Val: root}
	leftTree := inorder[:rootIndex]
	rightTree := inorder[rootIndex+1 : len(inorder)]

	if len(leftTree) > 0 {
		leftPreoder := preorder[1 : 1+len(leftTree)]
		rootNode.Left = buildTree(leftPreoder, leftTree)
	}
	if len(rightTree) > 0 {
		rightPreorder := preorder[1+len(leftTree):]
		rootNode.Right = buildTree(rightPreorder, rightTree)
	}

	return rootNode
}
