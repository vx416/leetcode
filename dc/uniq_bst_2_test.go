package dc

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestGenTrees(t *testing.T) {
	// t.Log(genNodes(3))
	t.Log(generateTrees(3))
}

// https://leetcode.com/problems/unique-binary-search-trees-ii/
func generateTrees(n int) []*TreeNode {
	nodes := genNodes(n)
	return genTrees(nodes)
}

func genTrees(nodes []*TreeNode) []*TreeNode {
	if len(nodes) == 1 {
		return nodes
	}

	trees := make([]*TreeNode, 0, len(nodes))

	for _, node := range nodes {
		lNodes, RNodes := seperateNodes(node, nodes)

		lTree := genTrees(lNodes)
		rTree := genTrees(RNodes)

		if len(lTree) == 0 {
			for _, rRoot := range rTree {
				root := clone(node)
				root.Right = rRoot
				trees = append(trees, root)
			}
		} else if len(rTree) == 0 {
			for _, lRoot := range lTree {
				root := clone(node)
				root.Left = lRoot
				trees = append(trees, root)
			}
		} else {
			for _, rRoot := range rTree {
				for _, lRoot := range lTree {
					root := clone(node)
					root.Left = lRoot
					root.Right = rRoot
					trees = append(trees, root)
				}
			}
		}
	}

	return trees
}

func clone(node *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   node.Val,
		Left:  node.Left,
		Right: node.Right,
	}
}

func genNodes(n int) []*TreeNode {
	nodes := make([]*TreeNode, 0, n)
	for i := 1; i <= n; i++ {
		node := &TreeNode{
			Val: i, Left: nil, Right: nil,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func seperateNodes(node *TreeNode, nodes []*TreeNode) ([]*TreeNode, []*TreeNode) {
	lNodes := make([]*TreeNode, 0, len(nodes)/2)
	rNodes := make([]*TreeNode, 0, len(nodes)/2)

	for _, n := range nodes {
		newN := clone(n)
		if newN.Val > node.Val {
			rNodes = append(rNodes, newN)
		} else if newN.Val < node.Val {
			lNodes = append(lNodes, newN)
		}
	}

	return lNodes, rNodes
}
