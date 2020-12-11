package tree

import (
	"testing"
)

func TestParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
	t.Log(generateParenthesis(1))
}

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	res := make([]string, 0, 1)
	parethesis("(", 1, 0, n, &res)
	return res
}

func parethesis(p string, left int, right int, target int, res *[]string) {
	if right == target && left == target {
		*res = append(*res, p)
		return
	}
	if left > target {
		return
	}
	if right > target {
		return
	}

	if p[len(p)-1] == '(' {
		parethesis(p+"(", left+1, right, target, res)
		parethesis(p+")", left, right+1, target, res)
	} else if p[len(p)-1] == ')' {
		parethesis(p+"(", left+1, right, target, res)
		if right < left {
			parethesis(p+")", left, right+1, target, res)
		}
	}
}
