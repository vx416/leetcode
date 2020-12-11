package recursion

import (
	"math"
	"testing"
)

func TestGrayCode(t *testing.T) {
	t.Log(grayCode(3))
}

// https://leetcode.com/problems/gray-code/
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	subAns := grayCode(n - 1)
	ans := make([]int, 0, len(subAns)*2)

	ans = append(ans, subAns...)

	for i := len(subAns) - 1; i >= 0; i-- {
		w := int(math.Pow(float64(2), float64(n-1)))
		ans = append(ans, subAns[i]+w)
	}

	return ans
}
