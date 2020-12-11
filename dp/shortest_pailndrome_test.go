package dp

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/shortest-palindrome/
func shortestPalindrome(s string) string {
	reS := reverse(s)
	if reS == s {
		return s
	}

	dp := prefixSuffix(s, reS)
	prefixStr := ""
	for i := len(s) - 1; i > 0; i-- {
		prefixStr += string(s[i])
		if dp[i] == i {
			break
		}
	}
	res := prefixStr + s

	if len(res)%2 == 1 {
		mid := len(res) / 2
		if res[mid] == res[mid+1] {
			return res[:mid] + res[mid+1:]
		}
	}
	return res

}

func reverse(s string) string {
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

// aacecaaa

func TestPrefixSuffix(t *testing.T) {
	// t.Log(reverse("abbac"))
	// t.Log(prefixSuffix("aabxbaac", reverse("aabxbaac")))
	t.Log(prefixSuffix("aacecaaa", reverse("aacecaaa")))
	// t.Log(shortestPalindrome("abbac"))
	t.Log(shortestPalindrome("aacecaaa"))
	t.Log(shortestPalindrome("aabba"))

	// t.Log(prefixSuffix("abbac", "cabba"))
}

func prefixSuffix(s string, rev string) []int {
	dp := make([]int, len(s))
	n := len(rev) - 1
	dp[0] = 0
	dp[1] = 1

	if s[0] == s[1] {
		dp[2] = 1
	}

	for i := 3; i < len(s); i++ {
		j := n - i
		// fmt.Printf("i:%d, j:%d\n", i, j)
		dp[i] = cal(s, rev, 0, j+1)
	}

	return dp
}

func cal(s, rev string, i, j int) int {
	res := 0

	for i < len(s) && j < len(rev) {
		fmt.Printf("i:%s, j:%s\n", string(s[i]), string(rev[j]))
		if s[i] == rev[j] {
			res++
			i++
			j++
		} else {
			break
		}
	}
	return res
}
