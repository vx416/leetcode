package dp

import (
	"fmt"
	"strconv"
)

// [["1","0","1","0","0","1","1","1","0"],
//  ["1","1","1","0","0","0","0","0","1"],
//  ["0","0","1","1","0","0","0","1","1"],
//  ["0","1","1","0","0","1","0","0","1"],
//  ["1","1","0","1","1","0","0","1","0"],
//  ["0","1","1","1","1","1","1","0","1"],
//  ["1","0","1","1","1","0","0","1","0"],
//  ["1","1","1","0","1","0","0","0","1"],
//  ["0","1","1","1","1","0","0","1","0"],
//  ["1","0","0","1","1","1","0","0","0"]]

// [[1 0 1 0 0 1 1 1 0]
//  [1 1 1 0 0 0 0 0 1]
//  [0 0 1 1 0 0 0 1 1]
//  [0 1 1 1 0 1 0 0 1]
//  [1 1 1 1 1 0 0 1 0]
//  [0 1 1 2 2 1 1 0 1]
//  [1 0 1 2 3 1 0 1 0]
//  [1 1 1 1 1 0 0 0 1]
//  [0 1 2 1 2 0 0 1 0]
//  [1 0 0 1 2 1 0 0 0]]

type square struct {
	size    int
	maxUp   int
	maxLeft int
}

func (s square) String() string {
	return fmt.Sprintf("%d", s.size)
}

// https://leetcode.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	intMatrix := make([][]*square, len(matrix))

	for i := range matrix {
		intMatrix[i] = make([]*square, len(matrix[i]))
		for j := range matrix[i] {
			num, _ := strconv.Atoi(string(matrix[i][j]))
			intMatrix[i][j] = &square{size: num, maxLeft: num, maxUp: num}
		}
	}

	rows := len(matrix)
	cols := len(matrix[0])
	max := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if intMatrix[r][c].size > max {
				max = intMatrix[r][c].size
			}
			if r-1 < 0 || c-1 < 0 {
				continue
			}
			if intMatrix[r][c].size > 0 {
				intMatrix[r][c].maxLeft = intMatrix[r][c-1].maxLeft + 1
				intMatrix[r][c].maxUp = intMatrix[r-1][c].maxUp + 1
				if intMatrix[r-1][c-1].size > 0 {
					intMatrix[r][c].size = min(intMatrix[r][c].maxUp, intMatrix[r][c].maxLeft, intMatrix[r-1][c-1].size+1)
					if intMatrix[r][c].size*intMatrix[r][c].size > max {
						max = intMatrix[r][c].size * intMatrix[r][c].size
					}
				}
			}

		}
	}

	return max
}

func min(nums ...int) int {
	min := 2 << 32
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func checkSize(matrix [][]byte, r, c, maxN int) int {
	newR := r
	newC := c
	n := 0
	for i := 1; i <= maxN; i++ {
		newR--
		newC--
		if newR < 0 || newC < 0 {
			break
		}

		if matrix[newR][c] == '0' {
			break
		}
		if matrix[r][newC] == '0' {
			break
		}
		n++
	}

	return n
}
