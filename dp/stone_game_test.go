package dp

import (
	"testing"
)

func TestStone(t *testing.T) {
	t.Log(stoneGame([]int{7, 8, 8, 10}))
}

// https://leetcode.com/problems/stone-game/
// dp[n, a, p] n => piles, a => alex, p => take tail or head
// dp[n, a, l] = max(dp[n+1, a, r], dp[n+1,a,l])) + dp[l]
// if dp[n,a,l][0] > dp[n,a,l][1] true else
// if dp[n,a,r][0] > dp[n,a,r][1] true else
// return false
// dp[i,j] = max(piles[i] + dp[i+1,j], piles[j] + dp[i, j-1])
func stoneGame(piles []int) bool {
	n := len(piles) - 1
	round := len(piles) / 2
	dp := make([]int, round)
	for i := range dp {
		head, tail := piles[i], piles[n-i]
		dp[i] = max(head, tail) - min(head, tail)
		if i > 0 {
			dp[i] += dp[i-1]
		}
	}
	return dp[round-1] > 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
