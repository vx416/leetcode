package stack

import (
	"testing"
)

func TestValidateStack(t *testing.T) {
	// t.Log(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	t.Log(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

// https://leetcode.com/problems/validate-stack-sequences/
func validateStackSequences(pushed []int, popped []int) bool {
	var (
		s             = stack{arr: make([]int, 0, 1)}
		push_i, pop_i int
	)

	for push_i < len(pushed) {
		// if can_pop{}
		popItem := popped[pop_i]
		for s.len() > 0 && s.top() == popItem {
			s.pop()
			pop_i++
			if pop_i < len(popped) {
				popItem = popped[pop_i]
			}
		}

		pushItem := pushed[push_i]
		s.push(pushItem)
		push_i++
	}

	if pop_i < len(popped) {
		popItem := popped[pop_i]
		for s.len() > 0 && s.top() == popItem {
			s.pop()
			pop_i++
			if pop_i < len(popped) {
				popItem = popped[pop_i]
			}
		}
	}
	return pop_i == len(popped)
}
