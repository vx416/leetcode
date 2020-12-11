package array

import "testing"

func TestMaxArea(t *testing.T) {
	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var (
		r    = 0
		l    = len(height) - 1
		area = 0
	)

	for r < l {
		if newArea := calArea(height, r, l); newArea > area {
			area = newArea
		}

		if height[r] > height[l] {
			l--
		} else {
			r++
		}
	}
	return area
}

func calArea(height []int, r, l int) int {
	var (
		w = l - r
		h = 0
	)
	if height[r] > height[l] {
		h = height[l]
	} else {
		h = height[r]
	}
	return w * h
}
