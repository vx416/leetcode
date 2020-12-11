package array

import (
	"testing"
)

func TestCircular(t *testing.T) {
	// t.Log(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	// t.Log(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	// t.Log(circularArrayLoop([]int{-1, 2}))
	// t.Log(circularArrayLoop([]int{-1, -1, -1, -1}))
	t.Log(circularArrayLoop([]int{-8, -1, 1, 7, 2}))
	t.Log(circularArrayLoop([]int{1, 1}))

}

func TestFixI(t *testing.T) {
	t.Log(fixIndex(3, -7))
	t.Log(fixIndex(2, 2))
}

func fixIndex(n int, i int) int {
	if i >= n {
		return i % n
	}

	if i < 0 {
		x := -i
		if x*n == 0 {
			return n - 1
		}
		return n - (x % n)
	}
	return i
}

//https://leetcode.com/problems/circular-array-loop/
func circularArrayLoop(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	flag := -1001
	for i := range nums {
		curr := i
		if nums[i] < -1000 {
			continue
		}

		for {
			nextI := curr + nums[curr]
			nextI = fixIndex(len(nums), nextI)
			if nextI == curr {
				break
			}
			if nums[nextI] == flag {
				return true
			}
			if nums[nextI] < -1000 && nums[nextI] != flag {
				break
			}
			if nums[nextI] < 0 && nums[curr] > 0 {
				break
			}
			if nums[nextI] > 0 && nums[curr] < 0 {
				break
			}
			nums[curr] = flag
			curr = nextI
		}
		flag--
	}

	return false
}
