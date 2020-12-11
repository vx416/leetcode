package array

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSumCloset(t *testing.T) {
	t.Log(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

// https://leetcode.com/problems/3sum-clos
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	res := target + 2<<32

	for i, num := range nums {
		towSum := twoSumClosest(nums[i+1:], target-num)
		fmt.Printf("num:%d, sum:%d\n", num, towSum)

		if absInt(towSum+num-target) < absInt(res-target) {
			res = towSum + num
		}

	}

	return res
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func twoSumClosest(nums []int, target int) int {
	i := 0
	j := len(nums) - 1

	res := target + 2<<32

	if nums[i] >= 0 {
	}

	for j >= 0 && i < len(nums) && nums[i] <= 0 && nums[j] > 0 {
		sum := nums[i] + nums[j]
		if sum == target {
			return sum
		}

		if absInt(sum-target) < absInt(res-target) {
			res = sum
		}
		if sum > target {
			j--
		} else if sum < target {
			i++
		}
	}

	return res
}

func binarySearch(nums []int, target int) int {
	i := 0
	j := len(nums)

	for i 

}
