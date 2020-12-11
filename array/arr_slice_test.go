package array

import (
	"testing"
)

// arr = [1,3,4,-1,3,5]
// 連續的 array arr_res 總合等於 n
// 找出 最大的 len(arr_res) arr_res
func TestArrSlice(t *testing.T) {
	t.Log(findMaxMatch([]int{1, 2, 4, -5, 3, 5, 1, 1, 1, 4, 1, 1}, 11))
}

func findMaxMatch(arr []int, n int) []int {
	dp := make([][]slice, len(arr))
	dp[0] = []slice{newSlice().add(arr[0])}
	res := []int{}

	for i := 1; i < len(arr); i++ {
		dp[i] = make([]slice, 0, 5)
		dp[i] = append(dp[i], newSlice().add(arr[i]))
		for _, preSlice := range dp[i-1] {
			newSlice := preSlice.add(arr[i])
			if newSlice.sum > n {
				continue
			}
			if newSlice.sum == n && newSlice.len() > len(res) {
				res = newSlice.arr
			}
			dp[i] = append(dp[i], newSlice)
		}
	}

	return res
}

type slice struct {
	arr []int
	sum int
}

func (s slice) len() int {
	return len(s.arr)
}

func (s slice) add(i int) slice {
	return slice{
		arr: append(s.arr, i),
		sum: s.sum + i,
	}
}

func newSlice() slice {
	return slice{
		arr: make([]int, 0, 1),
	}
}
