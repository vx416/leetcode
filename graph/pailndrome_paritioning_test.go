package graph

import (
	"testing"
)

func TestPartition(t *testing.T) {
	t.Log(partition("fffffffffffffff"))
	// t.Log(isPalindrome("bbabb"))
}

// https://leetcode.com/problems/palindrome-partitioning/
func partition(s string) [][]string {
	dp := make(map[string]bool)
	return findPartitions(s, dp)
}

func findPartitions(s string, dp map[string]bool) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}

	// if _, ok := dp[s]; ok {
	// 	return dp[s]
	// }

	res := make([][]string, 0, 5)
	if isPalindrome(s, dp) {
		res = append(res, []string{s})
	}

	for i := 0; i < len(s)-1; i++ {
		if isPalindrome(s[:i+1], dp) {
			rightPartitions := findPartitions(s[i+1:], dp)
			for _, p := range rightPartitions {
				res = append(res, append([]string{s[:i+1]}, p...))
			}
		}
	}
	// dp[s] = res
	return res
}

func isPalindrome(s string, dp map[string]bool) bool {
	if p, ok := dp[s]; ok {
		return p
	}

	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			dp[s] = false
			return false
		}
		i++
		j--
	}

	dp[s] = true

	return true
}
