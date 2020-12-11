package recursion

import "testing"

func TestPermutation(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}

// https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := make([][]int, 0, 1)
	for i := range nums {
		subNums := make([]int, 0, 1)
		subNums = append(subNums, nums[:i]...)
		subNums = append(subNums, nums[i+1:]...)
		subPermutation := permute(subNums)
		for j := range subPermutation {
			res = append(res, append([]int{nums[i]}, subPermutation[j]...))
		}
	}
	return res
}
