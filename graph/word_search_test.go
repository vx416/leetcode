package graph

import (
	"testing"
)

func TestWordSearch(t *testing.T) {
	// t.Log(exist([][]byte{{'A', 'B', 'C', 'E'}}, "ABCB"))
	// t.Log(exist([][]byte{{'a', 'b'}}, "ba"))
	// t.Log(exist([][]byte{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}}, "aaaaaaaaaaaaa"))
	t.Log(exist([][]byte{{'c', 'a', 'a'}, {'a', 'a', 'a'}, {'b', 'c', 'd'}}, "aab"))
	// t.Log(valid([][]byte{{'A', 'B', 'C', 'E'}}, [2]int{-1, 0}))
}

// https://leetcode.com/problems/word-search/
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				res := wordDfs(board, word, nil, []int{i, j}, 1)
				if res {
					return true
				}
			}
		}
	}
	return false
}

func wordDfs(board [][]byte, word string, prev, curr []int, i int) bool {
	if i >= len(word) {
		return true
	}
	orig := board[curr[0]][curr[1]]
	board[curr[0]][curr[1]] = '0'
	neighbors := getNeighor(board, prev, curr, rune(word[i]))
	res := false
	if len(neighbors) > 0 && i == len(word)-1 {
		return true
	}

	for _, n := range neighbors {
		res = wordDfs(board, word, curr, n, i+1) || res
	}
	board[curr[0]][curr[1]] = orig
	return res
}

func getNeighor(board [][]byte, prev, pos []int, target rune) [][]int {
	poss := make([][]int, 0, 1)

	// up
	up := []int{pos[0] - 1, pos[1]}
	if valid(board, up, target) && !eq(up, prev) {
		poss = append(poss, up)
	}

	// down
	down := []int{pos[0] + 1, pos[1]}
	if valid(board, down, target) && !eq(down, prev) {
		poss = append(poss, down)
	}
	// left
	left := []int{pos[0], pos[1] - 1}
	if valid(board, left, target) && !eq(left, prev) {
		poss = append(poss, left)
	}
	// right
	right := []int{pos[0], pos[1] + 1}
	if valid(board, right, target) && !eq(right, prev) {
		poss = append(poss, right)
	}

	return poss
}

func valid(board [][]byte, pos []int, target rune) bool {
	if pos[0] < 0 || pos[0] >= len(board) {
		return false
	}

	if pos[1] < 0 || pos[1] >= len(board[0]) {
		return false
	}

	if rune(board[pos[0]][pos[1]]) == target {
		return true
	}

	if rune(board[pos[0]][pos[1]]) == '0' {
		return false
	}

	return false
}

func eq(pos, pos1 []int) bool {
	if pos == nil || pos1 == nil {
		return false
	}
	return pos[0] == pos1[0] && pos[1] == pos1[1]
}
