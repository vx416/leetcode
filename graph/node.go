package graph

import (
	"strconv"
)

func arrToTree(arr []int) *TreeNode {
	nodes := make([]*TreeNode, len(arr)+1)

	for i := range arr {
		if arr[i] == -1 {
			nodes[i+1] = nil
		} else {
			nodes[i+1] = &TreeNode{Val: arr[i]}
		}
	}

	for i := range arr {
		if (i+1)*2 < len(nodes) {
			nodes[i+1].Left = nodes[(i+1)*2]
		}
		if (i+1)*2+1 < len(nodes) {
			nodes[i+1].Right = nodes[(i+1)*2+1]
		}
	}

	return nodes[1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	return strconv.Itoa(node.Val)
}
