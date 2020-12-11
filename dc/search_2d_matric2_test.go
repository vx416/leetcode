package dc

import (
	"testing"
)

func TestSearchM(t *testing.T) {
	matrix3 := [][]int{
		{3, 3, 8, 13, 13, 18},
		{4, 5, 11, 13, 18, 20},
		{9, 9, 14, 15, 23, 23},
		{13, 18, 22, 22, 25, 27},
		{18, 22, 23, 28, 30, 33},
		{21, 25, 28, 30, 35, 35},
		{24, 25, 33, 36, 37, 40},
	}
	t.Log(searchMatrix(matrix3, 21))
}

// https://leetcode.com/problems/search-a-2d-matrix-ii/
func searchMatrix(matrix [][]int, target int) bool {
	// for i := 0; i < len(matrix); i++ {
	// 	if binarySearch(matrix[i], target) {
	// 		return true
	// 	}
	// }
	// return false
	if len(matrix) == 0 {
		return false
	}
	return searchMatrix2(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, target)
}

func binarySearch(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return arr[0] == target
	}

	mid := len(arr) / 2
	if arr[mid] == target {
		return true
	}

	if arr[mid] > target {
		return binarySearch(arr[:mid], target)
	}

	return binarySearch(arr[mid+1:], target)
}

func searchMatrix2(matrix [][]int, rm, rM, cm, cM, target int) bool {
	if rm > rM || cm > cM {
		return false
	}

	if rm == rM && cm == cM {
		return matrix[rm][cm] == target
	}

	midR, midC := rm+(rM-rm)/2, cm+(cM-cm)/2
	if matrix[midR][midC] == target {
		return true
	}
	if target > matrix[rM][midC] && target > matrix[midR][cM] {
		return searchMatrix2(matrix, midR+1, rM, midC+1, cM, target)
	}

	if target < matrix[rm][midC] && target < matrix[midR][cm] {
		return searchMatrix2(matrix, rm, midR-1, cm, midC-1, target)
	}

	if target > matrix[midR][midC] {
		return searchMatrix2(matrix, rm, rM, midC+1, cM, target) || searchMatrix2(matrix, midR+1, rM, cm, midC, target)
	}

	return searchMatrix2(matrix, rm, midR-1, cm, cM, target) || searchMatrix2(matrix, midR, rM, cm, midC-1, target)
}
