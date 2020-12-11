package array

import (
	"testing"
)

func TestPurchase(t *testing.T) {
	t.Log(stocks([]int{1, 3, 11, 3, 6, 7, 40, 2, 5, 20, 10, 9, 8, 2, 10, 30, 50, 60}))
}

func stocks(arr []int) []int {
	var (
		min = arr[0]
		max = 0
		res = []int{}
	)

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if arr[i] > max {
				max = arr[i]
				if len(res) == 0 {
					res = []int{min, max}
				} else if max-min > res[1]-res[0] {
					res = []int{min, max}
				}
			}
		} else if arr[i] < arr[i-1] {
			if arr[i] < min {
				min = arr[i]
				max = 0
			}
		}
	}

	return res
}
