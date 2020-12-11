package tree

// https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	total := 0
	for i := range grid {
		for j := range grid[i] {
			node := [2]int{i, j}
			if grid[i][j] == '1' {
				bfs(grid, node)
				total++
			}
		}
	}

	return total
}

func bfs(grid [][]byte, node [2]int) {
	r, c := node[0], node[1]
	grid[r][c] = '2'
	nodes := getNeighbor(grid, node)
	for _, node := range nodes {
		if grid[node[0]][node[1]] == '1' {
			bfs(grid, node)
		}
	}
}

func getNeighbor(grid [][]byte, pos [2]int) [][2]int {
	neighbors := make([][2]int, 0, 1)

	rows := len(grid)
	cols := len(grid[0])

	// up
	rUP := pos[0] - 1
	cUP := pos[1]
	if rUP >= 0 && grid[rUP][cUP] == '1' {
		neighbors = append(neighbors, [2]int{rUP, cUP})
	}

	// down
	rDown := pos[0] + 1
	cDown := pos[1]
	if rDown < rows && grid[rDown][cDown] == '1' {
		neighbors = append(neighbors, [2]int{rDown, cDown})
	}

	// left
	rLeft := pos[0]
	cLeft := pos[1] - 1
	if cLeft >= 0 && grid[rLeft][cLeft] == '1' {
		neighbors = append(neighbors, [2]int{rLeft, cLeft})
	}

	// right
	rRight := pos[0]
	cRight := pos[1] + 1
	if cRight < cols && grid[rRight][cRight] == '1' {
		neighbors = append(neighbors, [2]int{rRight, cRight})
	}
	return neighbors
}
