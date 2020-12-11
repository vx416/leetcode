package tree

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	currNum, res := 0, 0
	inorder(root, k, &currNum, &res)
	return res
}

func inorder(curr *TreeNode, targetNum int, currNum, res *int) {
	if curr == nil {
		return
	}

	inorder(curr.Left, targetNum, currNum, res)
	*currNum++
	if *currNum == targetNum {
		*res = curr.Val
		return
	}
	inorder(curr.Right, targetNum, currNum, res)
}
