package tree

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *Node) *Node {
	levelNodes := make(map[int]*Node)

	dfs(root, levelNodes, 1)
	return root
}

func dfs(curr *Node, levelNodes map[int]*Node, level int) {
	if curr == nil {
		return
	}

	dfs(curr.Left, levelNodes, level+1)
	node, ok := levelNodes[level]
	if ok && node != nil {
		node.Next = curr
	}
	levelNodes[level] = curr
	dfs(curr.Right, levelNodes, level+1)
}
