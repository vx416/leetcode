package dc

import (
	"fmt"
	"testing"
)

func TestSimpleSearchMatrix(t *testing.T) {
	m := [][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50},
	}
	m2 := [][]int{
		{1, 3},
	}
	m3 := [][]int{
		{1, 1}, {2, 2},
	}
	t.Log(simpleSearchMatrix(m, 1))
	t.Log(simpleSearchMatrix(m2, 3))
	t.Log(simpleSearchMatrix(m3, 2))
}

// 0,1,2,4
// 4,5,6,7
// 1~8, 7 / 4 = 2,  7 % cols
// 3*4
// https://leetcode.com/problems/search-a-2d-matrix/
func simpleSearchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	left := 0
	right := (rows * cols) - 1

	for left <= right {
		mid := (right-left)/2 + left
		if mid >= rows*cols || mid < 0 {
			return false
		}
		midVal := getVal(matrix, mid)
		if midVal == target {
			return true
		}
		if midVal > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func getVal(matrix [][]int, i int) int {
	cols := len(matrix[0])
	if cols == 1 {
		return matrix[i][0]
	}
	if len(matrix) == 1 {
		return matrix[0][i]
	}
	r := (i + 1) / cols
	c := (i + 1) % cols

	if c > 0 {
		c--
	} else if c == 0 {
		c = cols - 1
		r--
	}

	fmt.Printf("cols:%d\n", cols)
	fmt.Printf("i:%d,r:%d,c:%d\n", i+1, r, c)
	return matrix[r][c]
}
