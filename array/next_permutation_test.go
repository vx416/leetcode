package array

import (
	"testing"
)

func TestNextPerm(t *testing.T) {
	nums := []int{2, 3, 1}
	nextPermutation(nums)
	t.Log(nums)
}

// https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	min := nums[len(nums)-1]
	minIndex := len(nums) - 1
	max := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		currNum := nums[i]
		if currNum < min {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
			reverse(nums[i+1:])
			return
		}
		if currNum > max {
			max = currNum
			continue
		}
		if currNum < max {
			j := findMinMaxIndex(nums, i+1, currNum)
			if j == -1 {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			reverse(nums[i+1:])
			return
		}
	}

	reverse(nums)
	return
}

func reverse(nums []int) {
	l := 0
	r := len(nums) - 1

	for r-l >= 1 {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func findMinMaxIndex(nums []int, startIndex int, min int) (maxIndex int) {
	if len(nums) == 0 {
		return -1
	}

	minMax := -1
	maxIndex = -1

	for i := startIndex; i < len(nums); i++ {
		num := nums[i]
		if num > min {
			if minMax == -1 || num <= minMax {
				minMax = num
				maxIndex = i
			}
		}
	}

	return maxIndex
}
