package array

import (
	"testing"
)

var mod = 1000000000 + 7

func TestThreeSum(t *testing.T) {
	t.Log(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
	t.Log(threeSumMulti([]int{1, 1, 2, 2, 2, 2}, 5))

	t.Log(threeSumMulti([]int{0, 0, 0}, 0))
	t.Log(threeSumMulti([]int{1, 1, 1}, 3))
}

// https://leetcode.com/problems/3sum-with-multiplicity/
func threeSumMulti(A []int, target int) int {
	res := 0
	max := 0
	counts := make(map[int]int)
	for _, val := range A {
		counts[val]++
		if val > max {
			max = val
		}
	}

	for i := 0; i <= max; i++ {
		// all diff
		for j := i + 1; j <= max; j++ {
			k := target - i - j
			if j < k && k <= max {
				res += counts[i] * counts[j] * counts[k]
				res %= mod
			}
		}
		// i == j
		k := target - i*2
		if k <= max && i < k {
			total := (counts[i] * (counts[i] - 1) / 2) * counts[k]
			res = ((res % mod) + total) % mod
		}

		// j == k
		if (target-i)%2 == 0 {
			j := (target - i) / 2
			if j <= max && i < j {
				total := counts[i] * counts[j] * (counts[j] - 1) / 2
				res = ((res % mod) + total) % mod
			}
		}
	}

	// i==j==k
	if target%3 == 0 {
		i := target / 3
		if counts[i] >= 3 && i <= 100 {
			total := counts[i] * (counts[i] - 1) * (counts[i] - 2) / 6
			res = ((res % mod) + total) % mod
		}
	}

	return res
}
