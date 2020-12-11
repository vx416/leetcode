package tree

// https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	max := 1 << 32
	min := -max
	return validateBST(root, min, max)
}

func validateBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	res := false
	if root.Val < max && root.Val > min {
		res = true
	}

	res = res &&
		validateBST(root.Left, min, root.Val) &&
		validateBST(root.Right, root.Val, max)
	return res
}
