package slidewin

import (
	"testing"
)

type queue struct {
	pos        map[int]int
	s          string
	head, tail int
}

func (q *queue) add(i int) int {
	ru := int(q.s[i])

	// check if exist in current queue
	pos, ok := q.pos[ru]
	if ok && pos >= q.head {
		q.head = pos + 1
	}
	q.pos[ru] = i
	q.tail = i
	return q.tail - q.head + 1
}

func TestLongLenSubStr(t *testing.T) {
	// t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("abbacabefcabcedx"))
	// t.Log(lengthOfLongestSubstring("bbbbb"))

}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	max := 0
	head := 0
	tail := 0
	posMap := make(map[rune]int)

	for i, ch := range s {
		pos, ok := posMap[ch]
		if ok && pos >= head {
			head = pos + 1
		}
		posMap[ch] = i
		tail = i
		if tail-head+1 > max {
			max = tail - head + 1
		}
	}
	return max
}
