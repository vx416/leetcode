package stack

import (
	"testing"
)

func Test132(t *testing.T) {
	// t.Log(find132patternStack([]int{1, 2, 3, 4}))
	// t.Log(find132patternStack([]int{3, 1, 4, 2}))
	t.Log(find132patternStack([]int{2, 4, 3, 1}))
}

// https://leetcode.com/problems/132-pattern/
func find132pattern(nums []int) bool {
	numI := 2 << 32
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] < numI {
			numI = nums[j]
		}
		for k := j + 1; k < len(nums); k++ {
			if nums[k] < nums[j] && numI < nums[k] {
				return true
			}
		}
	}
	return false
}

// min[i] = min(min[i-1], num[i])
func find132patternStack(nums []int) bool {
	minNums := findMinNums(nums)
	s := make([]int, 0, 1)

	for j := len(nums) - 1; j >= 1; j-- {
		if len(s) == 0 {
			s = append(s, nums[j])
			continue
		}

		if len(s) > 0 {
			if nums[j] > minNums[j] {
				// num[k] < num[j]
				for len(s) > 0 && s[len(s)-1] < nums[j] {
					numK := s[len(s)-1]
					s = s[:len(s)-1]
					if numK > minNums[j] {
						return true
					}
				}
				s = append(s, nums[j])
			}
		}
	}

	return false
}

func findMinNums(nums []int) []int {
	minNums := make([]int, len(nums))
	minNums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < minNums[i-1] {
			minNums[i] = nums[i]
		} else {
			minNums[i] = minNums[i-1]
		}
	}
	return minNums
}
