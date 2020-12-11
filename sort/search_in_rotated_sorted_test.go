package sort

import (
	"fmt"
	"testing"
)

func TestRotatedSort(t *testing.T) {
	t.Log(findSortedNums([]int{4, 5, 6, 7, 0, 1, 2}))
	t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {

	sortedNums := findSortedNums(nums)

	preIndex := 0
	for i := range sortedNums {
		got := searchSorted(sortedNums[i], target)
		if got != -1 {
			fmt.Println(got)
			return preIndex + got
		}
		preIndex = preIndex + len(sortedNums[i])
	}

	return 0
}

func searchSorted(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right-left > 1 {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] < target {
			left = mid
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func findSortedNums(nums []int) [][]int {
	res := make([][]int, 0, 1)

	tempRes := make([]int, 0, 1)
	for i := range nums {
		if i == 0 {
			tempRes = append(tempRes, nums[i])
			continue
		}
		if nums[i] < nums[i-1] {
			res = append(res, tempRes)
			tempRes = []int{nums[i]}
		} else {
			tempRes = append(tempRes, nums[i])
		}
	}
	res = append(res, tempRes)
	return res
}
