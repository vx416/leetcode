package dp

import "testing"

func TestMinSteps(t *testing.T) {
	t.Log(minSteps(3))
}

// https://leetcode.com/problems/2-keys-keyboard/
// dp[n][copy] = min步數
// dp[n][n/2] = min(dp[n/2][]) + 2
// dp[n][1] = dp[n-1][1]+1
// dp[n][2] = dp[n-2][2]+1 ....
func minSteps(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make(map[int]map[int]int)
	initMap(dp, 1)
	dp[1][0] = 0
	dp[1][1] = 1
	initMap(dp, 2)
	dp[2][1] = 2
	dp[2][2] = 3
	initMap(dp, 3)
	dp[3][1] = 3
	dp[3][3] = 4

	for i := 3; i <= n; i++ {
		initMap(dp, i)
		for j := 1; j < i/2; j++ {
			preStep, ok := dp[i-j][j]
			if ok {
				dp[i][j] = preStep + 1
			}
		}
		if i%2 == 0 {
			dp[i][i/2] = minInMap(dp, i/2) + 2
		}
	}

	return minInMap(dp, n)
}

func minInMap(dp map[int]map[int]int, key int) int {
	minVal := 2 << 32
	for _, v := range dp[key] {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func initMap(dp map[int]map[int]int, key int) {
	if _, ok := dp[key]; !ok {
		dp[key] = make(map[int]int)
	}
}
