package array

import (
	"strconv"
	"testing"
)

func TestRanges(t *testing.T) {
	t.Log(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	t.Log(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))

}

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var (
		res   = []string{}
		start = strconv.Itoa(nums[0])
		end   = ""
	)

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] || nums[i] == nums[i-1] {
			end = strconv.Itoa(nums[i])
		} else {
			if end != "" {
				res = append(res, start+"->"+end)
			} else {
				res = append(res, start)
			}
			start = strconv.Itoa(nums[i])
			end = ""
		}
	}

	if start != "" {
		if end != "" {
			res = append(res, start+"->"+end)
		} else {
			res = append(res, start)
		}
	}
	return res
}
