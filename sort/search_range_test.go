package sort

import (
	"testing"
)

func TestLeft(t *testing.T) {
	// t.Log(left([]int{5, 7, 7, 8, 8, 10}, 8))
	// t.Log(right([]int{5, 7, 7, 8, 8, 10}, 8))
	// t.Log(searchRange([]int{1}, 1))
	t.Log(searchRange([]int{1, 3}, 1))
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {

	res := []int{}
	leftIndex := left(nums, target)
	rightIndex := right(nums, target)
	if nums[leftIndex] == target {
		res = append(res, leftIndex)
	} else if leftIndex+1 < len(nums) && nums[leftIndex+1] == target {
		res = append(res, leftIndex+1)
	}
	if nums[rightIndex] == target {
		res = append(res, rightIndex)
	} else if rightIndex-1 >= 0 && nums[rightIndex-1] == target {
		res = append(res, rightIndex-1)
	}
	if len(res) == 2 {
		return res
	}

	return []int{-1, -1}
}

func right(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right-left > 1 {
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] <= target {
			left = mid
		}
	}

	return right
}

func left(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right-left > 1 {
		mid := (right-left)/2 + left
		if nums[mid] >= target {
			right = mid
		}
		if nums[mid] < target {
			left = mid
		}
	}

	return left
}
