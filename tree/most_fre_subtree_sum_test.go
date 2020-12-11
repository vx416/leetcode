package tree

// https://leetcode.com/problems/most-frequent-subtree-sum/
func findFrequentTreeSum(root *TreeNode) []int {
	sumFre := make(map[int]int)
	maxFre := 0
	findSum(root, sumFre, &maxFre)

	res := make([]int, 0, len(sumFre))
	for k, v := range sumFre {
		if v == maxFre {
			res = append(res, k)
		}
	}

	return res
}

func findSum(root *TreeNode, sumFre map[int]int, maxFre *int) int {
	if root == nil {
		return 0
	}

	leftSum := findSum(root.Left, sumFre, maxFre)
	rightSum := findSum(root.Right, sumFre, maxFre)

	treeSum := root.Val + leftSum + rightSum
	sumFre[treeSum]++
	if sumFre[treeSum] > *maxFre {
		*maxFre = sumFre[treeSum]
	}
	return treeSum
}
