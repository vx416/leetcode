package dp

import (
	"testing"
)

func TestSubSet(t *testing.T) {
	// t.Log(subsets([]int{1, 2, 3}))
	t.Log(subsets([]int{9, 0, 3, 5, 7}))
}

// https://leetcode.com/problems/subsets/
func subsets(nums []int) [][]int {
	sets := make([][]int, 0, 1)
	sets = append(sets, []int{})

	for _, num := range nums {

		tempSet := make([][]int, 0, 1)
		for _, set := range sets {
			tempSet = append(tempSet, set)
			newSet := make([]int, len(set))
			copy(newSet, set)
			newSet = append(newSet, num)
			tempSet = append(tempSet, newSet)
		}
		sets = tempSet
	}

	return sets
}
