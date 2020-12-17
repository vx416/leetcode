package tree

// https://leetcode.com/problems/flip-equivalent-binary-trees/
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil {
		return root1 == nil
	}

	if root1.Val != root2.Val {
		return false
	}

	if flipEquiv(root1.Left, root2.Left) {
		return flipEquiv(root1.Right, root2.Right)
	} else if flipEquiv(root1.Left, root2.Right) {
		return flipEquiv(root1.Right, root2.Left)
	}

	return false
}
