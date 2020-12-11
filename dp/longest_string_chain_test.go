package dp

import (
	"sort"
	"testing"
)

func TestLongestStrChain(t *testing.T) {
	// t.Log(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
	t.Log(longestStrChain([]string{"bdca", "bda", "ca", "dca", "a"}))
}

// https://leetcode.com/problems/longest-string-chain/
// dp[words] = max(dp[words-char[i]]+1)
func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}

	dp := make(map[string]int)
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	max := 1

	for i := range words {
		dp[words[i]] = 1
		if len(words) > 1 {
			subWords := breakWord(words[i])
			for _, subWord := range subWords {
				if n, ok := dp[subWord]; ok {
					if n+1 > dp[words[i]] {
						dp[words[i]] = n + 1
					}
				}
			}
		}
		if dp[words[i]] > max {
			max = dp[words[i]]
		}
	}

	return max
}

func breakWord(s string) []string {
	subStr := make([]string, 0, len(s))

	for i := range s {
		subStr = append(subStr, s[:i]+s[i+1:])
	}

	return subStr
}
