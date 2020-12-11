package dp

import "testing"

func TestSubSetWithDup(t *testing.T) {
	t.Log(subsetsWithDup([]int{1, 2, 2}))
}

// https://leetcode.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	setNums := make(map[int]int)
	for _, num := range nums {
		setNums[num]++
	}

	sets := make([][]int, 0, 1)
	sets = append(sets, []int{})

	for num, count := range setNums {
		tempSets := make([][]int, 0, 1)
		for _, set := range sets {
			tempSets = append(tempSets, set)
			for i := 1; i <= count; i++ {
				newSet := make([]int, len(set))
				copy(newSet, set)
				newSet = append(newSet, repeated(num, i)...)
				tempSets = append(tempSets, newSet)
			}
		}
		sets = tempSets
	}
	return sets
}

func repeated(v, n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = v
	}
	return res
}
