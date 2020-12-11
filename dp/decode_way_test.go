package dp

import (
	"strconv"
	"testing"
)

// 10, 1
// "10" => 2, 1, 10
// "1001" => 1,1, 10,1
// "1011"
// https://leetcode.com/problems/decode-ways/
// ans = dp[n, 2] + dp[n,1]
// if n <= 6
// dp[n, 2] = dp[n-2,2]+dp[n-2,1]
// dp[n,1] = dp[n-1,2]+dp[n-1,1]
func TestDecodeWay(t *testing.T) {
	// t.Log(numDecodings("226"))
	t.Log(numDecodings("1001"))
	// t.Log(numDecodings("11110"))
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][2]int, len(s))
	dp[0] = [2]int{1, 0}
	if s[0] == '0' {
		dp[0] = [2]int{0, 0}
	}

	for i := 1; i < len(s); i++ {
		// pick one
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		if s[i] == '0' {
			dp[i][0] = 0
		}
		// pick two
		if checkValid(s, i) {
			if i-2 < 0 {
				dp[i][1] = 1
			} else {
				dp[i][1] = dp[i-2][0] + dp[i-2][1]
			}
		} else {
			dp[i][1] = 0
		}
	}
	return dp[len(s)-1][0] + dp[len(s)-1][1]
}

func checkValid(s string, i int) bool {
	word := s[i-1 : i+1]
	num, _ := strconv.Atoi(word)
	return num <= 26 && num >= 10
}
