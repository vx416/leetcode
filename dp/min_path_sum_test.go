package dp

import "testing"

// https://leetcode.com/problems/minimum-path-sum/
// BFS => dp[i,j] = x[i,j] + min(dp[i,j-1], dp[i-1,j])
func TestMinPathSum(t *testing.T) {

	grid := [][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}
	t.Log(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	dp := initPathSumDP(rows, cols, grid)

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[rows-1][cols-1]
}

func initPathSumDP(rows, cols int, grid [][]int) [][]int {
	res := make([][]int, rows)
	for i := range res {
		res[i] = make([]int, cols)
	}
	res[0][0] = grid[0][0]

	for i := 1; i < rows; i++ {
		res[i][0] = grid[i][0] + res[i-1][0]
	}

	for j := 1; j < cols; j++ {
		res[0][j] = grid[0][j] + res[0][j-1]
	}
	return res
}
