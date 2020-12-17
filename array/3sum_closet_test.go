package array

import (
	"sort"
	"testing"
)

func TestThreeSumCloset(t *testing.T) {
	t.Log(threeSumClosest([]int{-4, -1, 1, 2}, 1))

	// nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	// t.Log(threeSumClosest(nums, -2))
	// // t.Log(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	// t.Log(bsClosest([]int{-5, -4, 0, 0, 3, 3, 4}, -2))
}

func twoSumCloset(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	closedNum := target + 1e9
	for l < r {
		twoSum := nums[l] + nums[r]
		if twoSum == target {
			return twoSum
		}
		if absDiff(twoSum, target) < absDiff(closedNum, target) {
			closedNum = twoSum
		}
		if twoSum < target {
			l++
		} else {
			r--
		}
	}

	return closedNum
}

// https://leetcode.com/problems/3sum-clos
func threeSumClosest(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	closedNum := target + 1e9
	for i := 0; i < len(nums)-2; i++ {
		twoSum := twoSumCloset(nums[i+1:], target-nums[i])
		if absDiff(target, twoSum+nums[i]) < absDiff(target, closedNum) {
			closedNum = twoSum + nums[i]
		}
	}

	return closedNum
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func bsClosest(nums []int, target int) (closet int) {
	l := 0
	r := len(nums) - 1

	for r-l > 1 {
		mid := (l + (r-l)/2)
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}

	if absDiff(target, nums[l]) > absDiff(target, nums[r]) {
		return nums[r]
	}
	return nums[l]
}
