package array

import (
	"fmt"
	"testing"
)

func TestTreeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// t.Log(threeSum([]int{0, 0, 0}))
	t.Log(threeSum([]int{3, -2, 1, 0}))
}

type set struct {
	arr [][]int
	s   map[string]bool
}

func (s *set) add(a ...int) {
	id := fmt.Sprintf("%d%d%d", a[0], a[1], a[2])
	if _, ok := s.s[id]; !ok {
		s.s[id] = true
		s.arr = append(s.arr, a)
	}
}

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	nums = quicksort(nums)

	res := &set{arr: make([][]int, 0, 1), s: make(map[string]bool)}

	i := 0
	if nums[i] == 0 {
		if nums[i+1] == 0 && nums[i+2] == 0 {
			return [][]int{{0, 0, 0}}
		}
		return [][]int{}
	}

	for nums[i] <= 0 {
		k := i + 1
		j := len(nums) - 1
		for k < j {
			sum := nums[i] + nums[k] + nums[j]
			if sum == 0 {
				res.add(nums[i], nums[k], nums[j])
				k++
				j--
			} else if sum < 0 {
				k++
			} else if sum > 0 {
				j--
			}
		}
		i++
		if nums[i] == nums[i-1] {
			i++
		}
	}

	return res.arr
}
