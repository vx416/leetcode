package sort

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	arr := []int{1, 0}
	// arr := []int{0, 1, 2, 1, 1, 3, 3, 2, 5}
	sortColors(arr)
	t.Log(arr)
}

// https://leetcode.com/problems/sort-colors/
func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}

	pivot := nums[0]
	i, j := 1, len(nums)-1
	for i < j {
		// 找到 i 的斷點 (nums[i]) > pivot
		for i < len(nums) && nums[i] <= pivot {
			i++
		}
		// 找到 j 的斷點 (nums[j]) < pivot
		for j >= 0 && nums[j] > pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// pivot is smallest
	if j == 0 {
		sortColors(nums[1:])
		return
	}

	// pivot is biggest
	if i == len(nums) {
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
		sortColors(nums[:len(nums)-1])
		return
	}

	// nums[i-1], pivot swap
	nums[0], nums[i-1] = nums[i-1], nums[0]
	sortColors(nums[:i-1])
	sortColors(nums[i:])
}
