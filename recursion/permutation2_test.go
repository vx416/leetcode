package recursion

import (
	"sort"
	"testing"
)

func TestPermutationUnique(t *testing.T) {
	t.Log(permuteUnique([]int{1, 1, 2}))

}

// https://leetcode.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)

	return permuteWithSorted(nums)
}

func permuteWithSorted(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := make([][]int, 0, 1)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		subNums := make([]int, 0, len(nums)-1)
		subNums = append(subNums, nums[:i]...)
		subNums = append(subNums, nums[i+1:]...)
		subPermutation := permuteWithSorted(subNums)
		for j := range subPermutation {
			res = append(res, append([]int{nums[i]}, subPermutation[j]...))
		}
	}
	return res
}
