package array

// https://leetcode.com/problems/set-matrix-zeroes/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	isCol := false

	for i := range matrix {
		if matrix[i][0] == 0 {
			isCol = true
		}

		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		matrix[0] = make([]int, len(matrix[0]))
	}

	if isCol {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
