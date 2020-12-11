package hash

import "testing"

// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
func TestGroupPepole(t *testing.T) {
	t.Log(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][][]int)

	init := func(i int) {
		if cap(m[i]) == 0 || m[i] == nil {
			m[i] = make([][]int, 0, 1)
		}
	}

	isFull := func(size int) bool {
		l := len(m[size])
		if l == 0 {
			return false
		}
		return len(m[size][l-1]) == size
	}

	add := func(size, i int) {
		l := len(m[size])
		if l == 0 {
			m[size] = append(m[size], []int{i})
		} else {
			m[size][l-1] = append(m[size][l-1], i)
		}
	}

	for i, size := range groupSizes {
		init(size)
		if isFull(size) {
			m[size] = append(m[size], []int{i})
		} else {
			add(size, i)
		}
	}

	result := make([][]int, 0, 1)

	for _, v := range m {
		result = append(result, v...)
	}

	return result
}
