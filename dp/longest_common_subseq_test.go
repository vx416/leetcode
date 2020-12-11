package dp

import "testing"

// https://leetcode.com/problems/longest-common-subsequence/

// dp[i,j] => text1[:i] and text2[:j] 的 longest common sub seq
// dp[i,j] = if text1[i] == text2[j]:  dp[i-1,j-1]+1
//   else text1[i] != text2[j]
//

func TestLongCommonSubSeq(t *testing.T) {
	t.Log(longestCommonSubsequence("abcde", "ace"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := initCommonSeqSeqDP(text1, text2)

	for i := 1; i < len(text2); i++ {
		for j := 1; j < len(text1); j++ {
			if text2[i] == text1[j] {
				dp[j][i] = dp[j-1][i-1] + 1
			} else {
				dp[j][i] = max(dp[j][i-1], dp[j-1][i])
			}
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}

func initCommonSeqSeqDP(text1, text2 string) [][]int {
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}

	for i := range text2 {
		if i > 0 {
			if dp[0][i-1] == 1 {
				dp[0][i] = 1
				continue
			}
		}
		if text1[0] == text2[i] {
			dp[0][i] = 1
		}
	}

	for i := range text1 {
		if i > 0 {
			if dp[i-1][0] == 1 {
				dp[i][0] = 1
				continue
			}
		}

		if text2[0] == text1[i] {
			dp[i][0] = 1
		}
	}
	return dp
}
