package sort

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	freq := make(map[int][]int)
	res := make([]int, 0, 1)

	for _, num := range nums {
		counts[num]++
	}
	max := 0
	for v, count := range counts {
		if count > max {
			max = count
		}

		if cap(freq[count]) == 0 {
			freq[count] = make([]int, 0, 1)
		}
		freq[count] = append(freq[count], v)
	}

	for i := max; i > 0; i-- {
		if items, ok := freq[i]; ok {
			res = append(res, items...)
		}
		if len(res) == k {
			break
		}
	}

	return res
}
