package array

import (
	"testing"
)

// https://leetcode.com/problems/maximize-distance-to-closest-person/
func TestMaxDistToCloset(t *testing.T) {
	t.Log(maxDistToClosest([]int{0, 1, 1, 1, 0, 0, 1, 0, 0}))
	// t.Log(maxDistToClosest([]int{0, 1, 0, 1, 0}))
}

func maxDistToClosest(seats []int) int {
	if len(seats) == 0 {
		return 0
	}

	i := 0
	// l := 0
	r := 0
	for i < len(seats) {
		if seats[i] == 1 {
			r = i
			i++
			break
		}
		i++
	}
	maxD := r - 0
	for i < len(seats) {
		if seats[i] == 1 {
			if (i-r)/2 > maxD {
				maxD = (i - r) / 2
			}
			r = i
		}
		i++
	}

	// fmt.Printf("l:%d,r:%d\n", l, r)

	if len(seats)-1-r > maxD {
		maxD = len(seats) - 1 - r
	}

	return maxD
}
