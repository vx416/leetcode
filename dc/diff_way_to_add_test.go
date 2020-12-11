package dc

import (
	"strconv"
	"testing"
)

func TestDiffWays(t *testing.T) {
	t.Log(diffWaysToCompute("2*3-4*5"))
}

// https://leetcode.com/problems/different-ways-to-add-parentheses/
func diffWaysToCompute(input string) []int {
	chars := getChars(input)
	return getWays(chars)
}

func getWays(s []Char) []int {
	if len(s) == 1 {
		return []int{s[0].Num}
	}

	var res = make([]int, 0, len(s))
	for i, x := range s {
		if x.IsOp() {
			lNums := getWays(s[:i])
			rNums := getWays(s[i+1:])
			for _, ln := range lNums {
				for _, rn := range rNums {
					res = append(res, x.Cal(ln, rn))
				}
			}
		}
	}
	return res
}

func getChars(s string) []Char {
	numChar := ""

	chars := make([]Char, 0, len(s))
	i := 0
	for i < len(s) {
		if isOp(string(s[i])) {
			if numChar != "" {
				num, _ := strconv.Atoi(numChar)
				chars = append(chars, Char{Num: num})
				numChar = ""
			}

			chars = append(chars, Char{Op: string(s[i])})
		} else {
			numChar = numChar + string(s[i])
		}
		i++
	}

	if numChar != "" {
		num, _ := strconv.Atoi(numChar)
		chars = append(chars, Char{Num: num})
	}

	return chars
}

type Char struct {
	Num int
	Op  string
}

func (c Char) IsOp() bool {
	return c.Op != ""
}

func (c Char) Cal(a, b int) int {
	switch c.Op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return 0
}

func isOp(s string) bool {
	return s == "*" || s == "+" || s == "-"
}
