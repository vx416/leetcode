package sort

import (
	"testing"
)

func TestPeak(t *testing.T) {
	// t.Log(findPeakElement([]int{1, 2, 1, 3, 5, 6, 7}))
	// t.Log(findPeakElement([]int{1, 2, 1, 3, 5, 6, 7, 3}))
	// t.Log(findPeakElement([]int{1, 2, 3, 4, 3}))
	// t.Log(findPeakElement([]int{3, 2, 1}))
	t.Log(findPeakElement([]int{-2147483648, -2147483647}))
}

// https://leetcode.com/problems/find-peak-element/
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for right-left+1 > 3 {
		mid := (right-left+1)/2 + left
		if mid == 0 {
			if nums[0] > nums[1] {
				return 0
			}
			left = 0
			continue
		}
		if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			}
			right = mid
			continue
		}

		if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
			return mid
		}
		if nums[mid-1] > nums[mid] {
			right = mid
		} else if nums[mid+1] > nums[mid] {
			left = mid
		}
	}

	if right-left+1 == 3 {
		return max(nums, left, left+1, right)
	}
	if right-left+1 == 2 {
		return max(nums, left, right)
	}

	return right
}

func max(nums []int, index ...int) int {
	max := -2 << 32
	maxI := 0
	for _, i := range index {
		if nums[i] > max {
			max = nums[i]
			maxI = i
		}
	}
	return maxI
}
