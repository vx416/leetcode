package slidewin

import (
	"testing"
)

func TestMaxSlide(t *testing.T) {
	// arr := &sortedArr{arr: []int{}, itemPos: make(map[int][]int)}
	// arr.insert(7)
	// arr.insert(2)
	// // arr.insert(3)
	// // arr.insert(5)
	// // arr.insert(5)
	// // arr.insert(4)
	// // arr.insert(100)
	// // arr.insert(50)
	// t.Logf("%+v", arr)
	// // arr.remove(50)
	// // arr.remove(4)
	// // arr.remove(5)
	// arr.remove(7)
	// t.Logf("%+v", arr)
	t.Log(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// t.Log(maxSlidingWindow([]int{1, 3, -1}, 1))
	// t.Log(maxSlidingWindow([]int{7, 2, 4}, 2))
}

// https://leetcode.com/problems/sliding-window-maximum/
// pop item is max
//
// pop item is not max
func maxSlidingWindow(nums []int, k int) []int {
	maxs := make([]int, 0, len(nums))

	q := &doubleQ{arr: make([]int, 0, k)}

	for i := range nums {
		if !q.empty() && i-q.head() == k {
			q.popHead()
		}

		// 移除所有小於 nums[i] 並且在 i 前面的 val
		for !q.empty() && nums[q.tail()] < nums[i] {
			q.popTail()
		}
		q.pushTail(i)

		if i >= k-1 {
			maxs = append(maxs, nums[q.head()])
		}
	}

	return maxs
}

type doubleQ struct {
	arr []int
}

func (q *doubleQ) pushTail(val int) {
	q.arr = append(q.arr, val)
}

func (q *doubleQ) empty() bool {
	return q.len() == 0
}
func (q *doubleQ) len() int {
	return len(q.arr)
}

func (q *doubleQ) head() int {
	return q.arr[0]
}

func (q *doubleQ) tail() int {
	return q.arr[len(q.arr)-1]
}

func (q *doubleQ) popHead() {
	q.arr = q.arr[1:]
}

func (q *doubleQ) popTail() {
	lastI := len(q.arr) - 1
	q.arr = q.arr[:lastI]
}
