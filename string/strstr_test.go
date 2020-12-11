package string

import "testing"

func TestStrStr(t *testing.T) {
	// 	"mississippi"
	// "pi"
	// "mississippi"
	// "issip"

	// "aabaaabaaac"
	// "aabaaac"
	// "adcadcaddcadde"
	// "adcadde"
	// t.Log(failedFunc2("abcacbabcacb"))
	// t.Log(failedFunc2("abcabcacab"))
	t.Log(getNext("aaabbaaaac"))
	t.Log(failFunc("aaabbaaaac"))
	t.Log(strStr("aabaaabaaac", "aabaaac"))
}

func strStr(haystack string, needle string) int {
	if haystack == "" && needle != "" {
		return -1
	}
	if haystack == "" && needle == "" {
		return 0
	}
	if haystack != "" && needle == "" {
		return 0
	}
	move := failedFunc(needle)

	i, j := 0, 0

	for i < len(haystack) && j < len(needle) {

		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = move[j]
			if j == -1 {
				i++
				j++
			}
		}
	}

	if j == len(needle) {
		return i - len(needle)
	}

	return -1
}

func failedFunc(s string) []int {
	dp := make([]int, len(s))
	dp[0] = -1
	if len(s) >= 2 {
		dp[1] = 0
	}

	for i := 2; i < len(s); i++ {
		dp[i] = 0
		if dp[i-1] > 0 {
			if s[i-1] == s[dp[i-1]] {
				dp[i] = dp[i-1] + 1
			}
		}
		if dp[i] == 0 && s[i-1] == s[0] {
			dp[i] = 1
		}
	}

	return dp
}

func failedFunc2(s string) []int {
	dp := make([]int, len(s))
	dp[0] = -1

	for j := 1; j < len(s); j++ {
		i := dp[j-1]
		for s[j] != s[i+1] && i >= 0 {
			i = dp[i]
		}

		if s[j] == s[i+1] {
			dp[j] = i + 1
		} else {
			dp[j] = -1
		}

	}
	return dp
}
