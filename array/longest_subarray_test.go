package array

import (
	"testing"
)

func TestLogSubArr(t *testing.T) {
	// t.Log(longestSubarray([]int{8, 2, 4, 7}, 4))
	// t.Log(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
	t.Log(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func longestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {
		return 0
	}

	queue := newMinMax(nums[0])
	maxLen := queue.len()
	for i := 1; i < len(nums); i++ {
		if queue.canAdd(nums[i], limit) {
			queue.add(nums[i])
			if queue.len() > maxLen {
				maxLen = queue.len()
			}
		} else {
			for {
				queue.remove()
				if queue.canAdd(nums[i], limit) {
					queue.add(nums[i])
					if queue.len() > maxLen {
						maxLen = queue.len()
					}
					break
				}
			}
		}
	}

	return maxLen
}

func newMinMax(i int) *minMaxSubArr {
	return &minMaxSubArr{
		arr: []int{i},
		min: i, max: i,
	}
}

type minMaxSubArr struct {
	arr      []int
	min, max int
}

func (arr minMaxSubArr) len() int {
	return len(arr.arr)
}

func (arr minMaxSubArr) canAdd(x, limit int) bool {
	if arr.len() == 0 {
		return true
	}

	if x < arr.min && arr.max-x > limit {
		return false
	}
	if x > arr.max && x-arr.min > limit {
		return false
	}
	return true
}

func (arr *minMaxSubArr) add(x int) minMaxSubArr {
	if arr.len() == 0 {
		arr.min = x
		arr.max = x
		arr.arr = []int{x}
		return *arr
	}

	min, max := arr.min, arr.max
	if x < min {
		min = x
	} else if x > max {
		max = x
	}

	arr.min = min
	arr.max = max
	arr.arr = append(arr.arr, x)
	return *arr
}

func (arr *minMaxSubArr) remove() {
	if arr.len() == 0 {
		return
	}

	if arr.len() == 1 {
		arr.min = 0
		arr.max = 0
		arr.arr = make([]int, 0, 1)
		return
	}

	x := arr.arr[0]
	arr.arr = arr.arr[1:]

	if x == arr.min {
		arr.resetMin()
	} else if x == arr.max {
		arr.resetMax()
	}
}

func (arr *minMaxSubArr) resetMin() {
	min := arr.arr[0]
	for _, val := range arr.arr {
		if val < min {
			min = val
		}
	}
	arr.min = min
}

func (arr *minMaxSubArr) resetMax() {
	max := 0
	for _, val := range arr.arr {
		if val > max {
			max = val
		}
	}
	arr.max = max
}
