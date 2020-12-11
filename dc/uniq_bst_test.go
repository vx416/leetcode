package dc

import "testing"

// https://leetcode.com/problems/unique-binary-search-trees/

func TestNumTrees(t *testing.T) {
	t.Log(numTrees(19))
}

func numTrees(n int) int {
	return getNumTree(1, n)
}

// dp[min, max, root] = dp[min, root-1, 1~root-1] * dp[root+1,max,root+1~max]

func getNumTree(min, max int) int {
	if min == max || min > max {
		return 1
	}

	var res int
	for i := min; i <= max; i++ {
		nRes := getNumTree(min, i-1) * getNumTree(i+1, max)
		res = res + nRes
	}

	return res
}

// func genInts(n int) []int {
// 	nodes := make([]int, 0, n)

// 	for i := 1; i <= n; i++ {
// 		nodes = append(nodes, i)
// 	}
// 	return nodes
// }

// func seperate(v int, nodes []int) ([]int, []int) {
// 	lNodes := make([]int, 0, len(nodes))
// 	rNodes := make([]int, 0, len(nodes))

// 	for _, n := range nodes {
// 		if n > v {
// 			rNodes = append(rNodes, n)
// 		} else if n < v {
// 			lNodes = append(lNodes, n)
// 		}
// 	}

// 	return lNodes, rNodes
// }
