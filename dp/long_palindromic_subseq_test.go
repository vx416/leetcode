package dp

import (
	"fmt"
	"testing"
)

func TestLongPailSubSeq(t *testing.T) {
	// t.Log(longestPalindromeSubseq("bbbab"))
	// t.Log(longestPalindromeSubseq("cbbd"))
	// t.Log(longestPalindromeSubseq("aa"))
	// t.Log(longestPalindromeSubseq("abcdef"))
	t.Log(longestPalindromeSubseq("abcabcabcabc"))
}

// dp[len,i,j]
// dp[i,j]
// dp[1][0,0] = 1
// dp[2][0,1] = 2
// dp[3][0,2] = if 0 == 2: dp[2][1,1]+2 else max(dp[2][1,2],dp[2][0,1])
// dp[4][0,3] if 0 != 3: dp[3][0,2],dp[3][1,3]

// https://leetcode.com/problems/longest-palindromic-subsequence/
//  dp[i-1,i+1] = if i-1==i+1 == dp[i][i]+2 else = dp[i][i]
//  dp[i,i+1]
//  dp[i-1,i]
func Range(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func longestPalindromeSubseq(s string) int {
	dp := make([]map[string]int, len(s)+1)

	for i := range dp {
		dp[i] = make(map[string]int)
	}

	for i := range s {
		// pos := 1 << i
		dp[1][Range(i, i)] = 1
	}

	for i := 0; i < len(s)-1; i++ {
		// pos := 1<<i | 1<<(i+1)
		if s[i] == s[i+1] {
			dp[2][Range(i, i+1)] = 2
			// dp[2][i][i+1] = 2
		} else {
			dp[2][Range(i, i+1)] = 1
			// dp[2][i][i+1] = 1
		}
	}

	// i 是長度
	for i := 3; i <= len(s); i++ {
		// j 是起點
		for j := 0; j < len(s)-i+1; j++ {
			end := j + i - 1
			if s[j] == s[end] {
				// prePos := 1<<(j+1) | 1<<(end-1)
				dp[i][Range(j, end)] = dp[i-2][Range(j+1, end-1)] + 2
				// dp[i][j][end] = dp[i-2][j+1][end-1] + 2
			} else {
				// prePos1 := 1<<j | 1<<(end-1)
				// prePos2 := 1<<(j+1) | 1<<end
				dp[i][Range(j, end)] = max(dp[i-1][Range(j, end-1)], dp[i-1][Range(j+1, end)])
				// dp[i][j][end] = max(dp[i-1][j][end-1], dp[i-1][j+1][end])
			}
		}
	}

	return dp[len(s)][Range(0, len(s)-1)]
}
