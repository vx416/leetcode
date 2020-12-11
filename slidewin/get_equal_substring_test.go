package slidewin

import (
	"testing"
)

// https://leetcode.com/problems/get-equal-substrings-within-budget/
// head tail if s[tail] == t[tail] return max
// else if cost < maxCost return max
// else remove head

type win struct {
	head, tail int
	// cost       []int
	maxCost  int
	currCost int
	s, t     string
}

func (w *win) len() int {
	if w.tail == -1 || w.head == -1 {
		return 0
	}

	return w.tail - w.head + 1
}

func (w *win) add(i int) int {
	newCost := abs(int(w.s[i]) - int(w.t[i]))
	if newCost > w.maxCost {
		w.head = -1
		w.tail = -1
		w.currCost = -1
		return w.len()
	}

	if w.len() == 0 && newCost <= w.maxCost {
		w.head = i
		w.tail = i
		w.currCost = newCost
		return w.len()
	}

	if w.currCost+newCost <= w.maxCost {
		w.tail = i
		w.currCost += newCost
		return w.len()
	}

	removeCost := w.currCost + newCost - w.maxCost
	newH, cost := w.newHead(removeCost)
	// remove all windows
	if newH >= w.tail {
		if newCost <= w.maxCost {
			w.tail = i
			w.head = i
			w.currCost = newCost
			return w.len()
		}
		w.tail = -1
		w.head = -1
		w.currCost = -1
		return w.len()
	}

	w.head = newH
	w.currCost = w.currCost - cost + newCost
	return w.len()
}

func (w *win) newHead(removeCost int) (int, int) {
	cost := 0
	i := w.head
	for cost < removeCost && i <= w.tail {
		cost += abs(int(w.s[i]) - int(w.t[i]))
		i++
	}
	return i, cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func TestEqualSubstring(t *testing.T) {
	// t.Log(equalSubstring("abcd", "bcdf", 3))
	// t.Log(equalSubstring("abcd", "cdef", 1))
	// t.Log(equalSubstring("abcd", "adef", 0))
	t.Log(equalSubstring("abcd", "abcd", 0))
	// t.Log(equalSubstring("iktqzyazth", "havakbjzzc", 78))
	// t.Log(equalSubstring("abcdefghij", "bcdefghijk", 5))
}

func equalSubstring(s string, t string, maxCost int) int {
	w := &win{
		s: s, t: t,
		maxCost: maxCost,
		tail:    -1, head: -1,
		currCost: -1,
	}

	max := 0
	for i := range s {
		l := w.add(i)
		if l > max {
			max = l
		}
	}

	return max
}
