package sort

import (
	"testing"
)

func TestFindKth(t *testing.T) {
	// t.Log(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	t.Log(findKthLargest([]int{-1, -1, -2, -3}, 3))
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	counts := make(map[int]int)
	max := -2 << 32
	min := 2 << 32
	res := 0

	for _, num := range nums {
		counts[num]++
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	currentOrder := 0
	for i := max; i >= min; i-- {
		count, ok := counts[i]
		if ok {
			currentOrder += count
			if currentOrder >= k {
				res = i
				break
			}
		}
	}
	return res
}
