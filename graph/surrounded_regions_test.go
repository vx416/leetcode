package graph

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	board1 := [][]byte{
		{'O', 'X', 'O', 'X'},
		{'O', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X'},
	}
	solve(board1)
	t.Log(board1)
}

// https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte) {
	visited := make(map[string]bool)
	for i := range board {
		for j := range board[i] {
			pos := [2]int{i, j}
			if board[i][j] == 'O' && !visited[key(pos)] {
				n := [][2]int{pos}
				if !dfs(board, pos, visited, &n) {
					for _, nPos := range n {
						board[nPos[0]][nPos[1]] = 'X'
					}
				}
			}
		}
	}
}

func dfs(board [][]byte, start [2]int, visited map[string]bool, n *[][2]int) bool {
	isBorder := false
	if start[0] == len(board)-1 || start[0] == 0 || start[1] == len(board[0])-1 || start[1] == 0 {
		isBorder = true
	}
	visited[key(start)] = true

	neighbors, hasBorder := findNeighbor(board, start[0], start[1])
	for _, neighbor := range neighbors {
		if !visited[key(neighbor)] {
			*n = append(*n, neighbor)
			isBorder = dfs(board, neighbor, visited, n) || isBorder
		}
	}
	return isBorder || hasBorder
}

func key(pos [2]int) string {
	return fmt.Sprintf("r:%d,c:%d", pos[0], pos[1])
}

func findNeighbor(board [][]byte, row, col int) ([][2]int, bool) {
	rowMax := len(board)
	colMax := len(board[0])
	neighbors := make([][2]int, 0, 4)
	hasBorder := false
	// go up
	if col-1 >= 0 && board[row][col-1] == 'O' {
		neighbors = append(neighbors, [2]int{row, col - 1})
		if col-1 == 0 {
			hasBorder = true
		}
	}
	// go down
	if col+1 < colMax && board[row][col+1] == 'O' {
		neighbors = append(neighbors, [2]int{row, col + 1})
		if col+1 == colMax-1 {
			hasBorder = true
		}
	}
	// go right
	if row+1 < rowMax && board[row+1][col] == 'O' {
		neighbors = append(neighbors, [2]int{row + 1, col})
		if row+1 == rowMax-1 {
			hasBorder = true
		}
	}
	// go left
	if row-1 >= 0 && board[row-1][col] == 'O' {
		neighbors = append(neighbors, [2]int{row - 1, col})
		if row-1 == 0 {
			hasBorder = true
		}
	}
	return neighbors, hasBorder
}
