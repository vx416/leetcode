package tree

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

type NodePath struct {
	*TreeNode
	Depth int
	Where string
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}
	res := make([]int, 0, 10)
	count(target, K, 0, &res)
	if root.Val == target.Val {
		return res
	}

	nodePaths, _ := findTarget(root, target.Val)
	for _, path := range nodePaths {
		if path.Depth == K {
			res = append(res, path.Val)
			continue
		}

		if path.Where == "left" {
			if path.Depth < K {
				count(path.TreeNode.Right, K-path.Depth, 1, &res)
			}
		} else if path.Where == "right" {
			if path.Depth < K {
				count(path.TreeNode.Left, K-path.Depth, 1, &res)
			}
		}
	}
	return res
}

func count(node *TreeNode, k, depth int, res *[]int) {
	if node == nil {
		return
	}
	if k == depth {
		*res = append(*res, node.Val)
	}
	count(node.Left, k, depth+1, res)
	count(node.Right, k, depth+1, res)
}

func findTarget(root *TreeNode, target int) ([]*NodePath, int) {
	if root == nil {
		return nil, 0
	}
	if root.Left != nil && root.Left.Val == target {
		return []*NodePath{&NodePath{TreeNode: root, Depth: 1, Where: "left"}}, 2
	}
	if root.Right != nil && root.Right.Val == target {
		return []*NodePath{&NodePath{TreeNode: root, Depth: 1, Where: "right"}}, 2
	}

	nodesPath, preDepth := findTarget(root.Left, target)
	if len(nodesPath) != 0 {
		nodesPath = append(nodesPath, &NodePath{TreeNode: root, Depth: preDepth, Where: "left"})
		return nodesPath, preDepth + 1
	}
	nodesPath, preDepth = findTarget(root.Right, target)
	if len(nodesPath) != 0 {
		nodesPath = append(nodesPath, &NodePath{TreeNode: root, Depth: preDepth, Where: "right"})
	}

	return nodesPath, preDepth + 1
}
