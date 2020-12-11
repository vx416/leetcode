package recursion

import "testing"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

var numDigits = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func TestLetterComb(t *testing.T) {
	t.Log(letterCombinations("2"))
}

func letterCombinations(digits string) []string {
	// return nil
	letter := make([][]string, 0, 1)
	for _, d := range digits {
		letter = append(letter, numDigits[d])
	}

	return genString(letter, "")
}

func genString(letterSlice [][]string, s string) []string {
	if len(letterSlice) == 0 {
		return []string{}
	}

	if len(letterSlice) == 1 {
		res := make([]string, 0, 1)
		for i := range letterSlice[0] {
			res = append(res, s+letterSlice[0][i])
		}
		return res
	}

	res := make([]string, 0, 1)
	for i := range letterSlice[0] {
		subRes := genString(letterSlice[1:], s+letterSlice[0][i])
		res = append(res, subRes...)
	}
	return res
}
