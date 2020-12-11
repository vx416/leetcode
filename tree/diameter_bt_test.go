package tree

//  https://leetcode.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	maxLen, _ := maxLen(root)
	return maxLen
}

func maxLen(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Right == nil && root.Left == nil {
		return 0, 1
	}

	leftMax, leftLen := maxLen(root.Left)
	rightMax, rightLen := maxLen(root.Right)
	currMax := leftLen + rightLen
	currMax = max(currMax, leftMax)
	currMax = max(currMax, rightMax)

	return currMax, max(leftLen, rightLen) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
