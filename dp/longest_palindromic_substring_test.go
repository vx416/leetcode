package dp

import (
	"testing"
)

func TestLP(t *testing.T) {
	// t.Log(longestPalindrome("babad"))
	// t.Log(longestPalindrome("cbbd"))
	// t.Log(longestPalindrome("ccc"))
	// t.Log(longestPalindrome("abb"))
	// t.Log(longestPalindrome("abcba"))
	// t.Log(longestPalindrome("aaabaaaa"))
	t.Log(longestPalindrome("jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"))
}

// https://leetcode.com/problems/longest-palindromic-substring/
// dp[i,j] = dp[i+1,j-1], s[i]==s[j],
func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}

	var max string

	for i := 0; i < len(s); i++ {
		p1 := findPalindrome(s, i, i)
		p2 := findPalindrome(s, i, i+1)
		if len(max) < len(p1) {
			max = p1
		}

		if len(max) < len(p2) {
			max = p2
		}
	}
	return max

}

func findPalindrome(s string, lower, upper int) string {
	for lower >= 0 && upper < len(s) && s[lower] == s[upper] {
		lower--
		upper++
	}

	if lower == upper {
		return ""
	}

	return s[lower+1 : upper]

}
