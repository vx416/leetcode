package array

import (
	"testing"
)

func TestIncresTrip(t *testing.T) {
	// t.Log(increasingTriplet([]int{1, 2, 3, 4, 5}))
	// t.Log(increasingTriplet([]int{5, 4, 3, 2, 1}))
	// t.Log(increasingTriplet([]int{1, 2, 3, 1, 2, 1}))
	t.Log(increasingTriplet([]int{5, 1, 5, 5, 2, 5, 4}))
}

// https://leetcode.com/problems/increasing-triplet-subsequence/
// i < j < k, nums[i] < nums[j] < nums[k]
// i ~ k,
//
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min1, min2 := 2<<32, 2<<32

	for i := 0; i < len(nums); i++ {
		if nums[i] < min1 {
			min1 = nums[i]
		} else if nums[i] < min2 { // min2 > min1
			min2 = nums[i]
		} else {
			return true
		}
	}

	return false
}

func increasingTriplet2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min := nums[0]
	max := min

	for k := 2; k < len(nums); k++ {
		if nums[k] > max {
			max = nums[k]
		}
	}

	for j := 1; j < len(nums)-1; j++ {
		if nums[j] > max {
			max = 0
			for k := 3; k < len(nums); k++ {
				if nums[k] > max {
					max = nums[k]
				}
			}
		}

		if nums[j] > min && nums[j] < max {
			return true
		}

		if nums[j] < min {
			min = nums[j]
		}
	}

	return false
}
