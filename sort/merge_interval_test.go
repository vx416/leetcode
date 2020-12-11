package sort

import (
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Log(merge([][]int{
		{1, 3}, {8, 10}, {2, 6}, {15, 18},
	}))
}

// https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	sorted := SortBy(intervals)
	sort.Sort(sorted)

	res := make([][]int, 0, 1)

	res = append(res, sorted[0])

	for i := 1; i < len(sorted); i++ {
		last := len(res) - 1
		merged := tryMerge(res[last], sorted[i])
		if len(merged) > 0 {
			res[last] = merged
		} else {
			res = append(res, sorted[i])
		}
	}

	return res
}

func tryMerge(x, y []int) []int {
	if !hasMerge(x, y) {
		return []int{}
	}
	return []int{getMin(x[0], y[0]), getMax(x[1], y[1])}
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func hasMerge(x, y []int) bool {
	if y[0] > x[1] {
		return false
	}
	return true
}

type SortBy [][]int

func (a SortBy) Len() int {
	return len(a)
}

func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a SortBy) Less(i, j int) bool {
	x := a[i]
	y := a[j]

	if x[0] == y[0] {
		return x[1] < y[1]
	}
	return x[0] < y[0]
}
