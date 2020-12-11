package dp

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/unique-paths-ii/
// grapth BFS
// 每掃過一次 grid(i,j) => grid(i,j) + grid(i-1,j) + grid(i,j-1)
func TestUniqPath2(t *testing.T) {
	test4 := [][]int{
		{0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 0, 0, 0, 0},
	}
	t.Log(uniquePathsWithObstacles(test4))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var ans int
	if len(obstacleGrid) == 0 {
		return ans
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	dp := initDP(rows, cols, obstacleGrid)

	stack := &Stack{
		ps:    make([]Point, 0, 1),
		exist: make(map[string]bool),
	}
	stack.Push(NewP(0, 0))
	BFS(G(obstacleGrid), stack, dp)
	return dp.Get()[rows-1][cols-1]
}

func NewP(x, y int) Point {
	return Point{
		X: x, Y: y,
	}
}

type Point struct {
	X, Y int
}

func (p Point) ID() string {
	return fmt.Sprintf("X:%d,Y:%d", p.X, p.Y)
}

func (p Point) InValid() bool {
	return p.X == 0 || p.Y == 0
}

func BFS(g G, stack *Stack, dp *DP) {
	parent := stack.Pop()
	adjacencyP := g.GetAdjacency(parent)
	vals := dp.Get()
	for _, p := range adjacencyP {
		if !p.InValid() {
			vals[p.X][p.Y] += vals[parent.X][parent.Y]
		}
		stack.Push(p)
	}

	dp.Set(vals)
	if stack.Empty() {
		return
	}
	BFS(g, stack, dp)
}

type Stack struct {
	ps    []Point
	exist map[string]bool
}

func (s *Stack) Push(p Point) {
	if s.exist[p.ID()] {
		return
	}

	s.ps = append(s.ps, p)
	s.exist[p.ID()] = true
}

func (s *Stack) Pop() Point {
	p := s.ps[0]
	s.ps = s.ps[1:]
	return p
}

func (s *Stack) Empty() bool {
	return len(s.ps) == 0
}

type G [][]int

func (g G) GetAdjacency(p Point) []Point {
	ps := make([]Point, 0, 2)

	rigthP := NewP(p.X+1, p.Y)
	if !g.HasObstacle(rigthP) {
		ps = append(ps, rigthP)
	}

	downP := NewP(p.X, p.Y+1)
	if !g.HasObstacle(downP) {
		ps = append(ps, downP)
	}

	return ps
}

func (g G) HasObstacle(p Point) bool {
	if p.X >= len(g) {
		return true
	}
	if p.Y >= len(g[0]) {
		return true
	}

	return g[p.X][p.Y] == 1
}

type DP struct {
	values [][]int
}

func (dp *DP) Get() [][]int {
	return dp.values
}

func (dp *DP) Set(values [][]int) {
	dp.values = values
}

func initDP(rows, cols int, grid [][]int) *DP {
	dp := make([][]int, rows)

	for i := range dp {
		dp[i] = make([]int, cols)
	}

	dp[0][0] = 1

	for i := 1; i < cols; i++ {
		if grid[0][i] != 1 && grid[0][i-1] != 1 {
			dp[0][i] = 1
		}
	}
	for i := 1; i < rows; i++ {
		if grid[i][0] != 1 && grid[i-1][0] != 1 {
			dp[i][0] = 1
		}
	}
	return &DP{values: dp}
}
