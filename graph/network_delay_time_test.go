package graph

import (
	"fmt"
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	times3 := [][]int{
		{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59}, {5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43},
		{4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16}, {3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0},
		{2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59},
	}
	t.Log(networkDelayTime(times3, 5, 5))

}

// dp[k, i] = min(dp[k, adjacent_nodes_of_i]) + distance(adjacent_nodes_of_i, i)
// when dp[k, node_i] is determined
// update all node_i's adjacent_nodes(j) if dp[k, node_j] > dp[k, node_i] + distance(node_i, node_j)

// notice: distance can be zero

// https://leetcode.com/problems/network-delay-time/
func networkDelayTime(times [][]int, N int, K int) int {
	kToNodeCost := make(map[int]int)
	minArrival := make(map[int]bool)
	minArrival[K] = true
	kToNodeCost[K] = 0
	currMinArrival := K
	g := makeGraph(times, N)

	for len(minArrival) != N {
		edges := g.FindNeighbor(currMinArrival)
		// fmt.Printf("curr:%d\n", currMinArrival)
		for _, e := range edges {
			_, ok := kToNodeCost[e.to]

			// fmt.Printf("%s\n", e)
			if !ok || kToNodeCost[currMinArrival]+e.cost < kToNodeCost[e.to] {
				kToNodeCost[e.to] = kToNodeCost[currMinArrival] + e.cost
			}
		}
		currMinArrival = nextMinArrival(kToNodeCost, minArrival)
		if currMinArrival == -1 {
			break
		}
		// fmt.Printf("next:%d\n", currMinArrival)
		minArrival[currMinArrival] = true
		// fmt.Printf("%+v\n", kToNodeCost)
	}

	if len(minArrival) != N {
		return -1
	}

	maxCost := 0
	for _, cost := range kToNodeCost {
		if cost > maxCost {
			maxCost = cost
		}
	}

	return maxCost
}

func nextMinArrival(kToNodeCost map[int]int, minArrival map[int]bool) int {
	minCost := int(1e9)
	newMinArrival := -1

	for node, cost := range kToNodeCost {
		if cost < minCost && !minArrival[node] {
			newMinArrival = node
			minCost = cost
		}
	}

	return newMinArrival
}

type Edge struct {
	from int
	to   int
	cost int
}

func (e Edge) String() string {
	return fmt.Sprintf("%d=>%d cost:%d", e.from, e.to, e.cost)
}

type Graph [][]int

func (g Graph) FindNeighbor(i int) []Edge {
	edges := make([]Edge, 0, len(g))
	for j, cost := range g[i] {
		if cost != -1 {
			e := Edge{
				from: i,
				to:   j,
				cost: cost,
			}
			edges = append(edges, e)
		}
	}

	return edges
}

func makeGraph(times [][]int, nodes int) Graph {
	g := make([][]int, nodes+1)
	for i := range g {
		g[i] = make([]int, nodes+1)
		for j := range g[i] {
			g[i][j] = -1
		}
	}

	for _, edge := range times {
		from := edge[0]
		to := edge[1]
		cost := edge[2]
		g[from][to] = cost
	}

	return g
}
