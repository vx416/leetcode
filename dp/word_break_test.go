package dp

import (
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.com/problems/word-break/

// "catsandogcat"
// ["cats","dog","sand","and","cat","an"]
func TestWordBreak(t *testing.T) {
	// t.Log(KMPSearch("catsan", "sand", move("sand")))
	// t.Log(KMPSearch("aabcd", "abc", move("abc")))
	// t.Log(wordBreak("applepenapple", []string{"apple", "pen"}))
	// t.Log(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// t.Log(wordBreak("abcd", []string{"a", "abc", "b", "cd"}))
	t.Log(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// dp[s-word] = true
func wordBreak(s string, wordDict []string) bool {
	dp := make(map[string]int)
	// moveMap := make(map[string][]int)

	for i := range wordDict {
		dp[wordDict[i]] = 1
		// moveMap[wordDict[i]] = move(wordDict[i])
	}

	return wordBreaks(s, wordDict, dp)
}

func wordBreaks(s string, words []string, dp map[string]int) bool {
	if len(s) == 0 {
		return false
	}

	if x, ok := dp[s]; ok && x == -1 {
		return false
	}

	if x, ok := dp[s]; ok && x == 1 {
		return true
	}

	for i := range words {
		if words[i] == s {
			return true
		}
		if len(words[i]) > len(s) {
			continue
		}

		breakIndex := strings.Index(s, words[i])
		if breakIndex == -1 {
			// dp[words[i]] = -1
			continue
		}

		aS := s[:breakIndex]
		aSFound := true
		fmt.Printf("a:%s, w:%s s:%s\n", aS, words[i], s)
		if aS != "" {
			if !wordBreaks(aS, words, dp) {
				aSFound = false
				dp[aS] = -1
			}
		}

		bS := s[breakIndex+len(words[i]):]
		bSFound := true
		if bS != "" {
			if !wordBreaks(bS, words, dp) {
				bSFound = false
				dp[bS] = -1
			}
		}
		if aSFound && bSFound {
			return true
		}

	}

	return false
}
