package sort

import (
	"fmt"
	"testing"
)

func TestActivityNotifications(t *testing.T) {
	res := activityNotifications([]int32{10, 20, 30, 40, 50}, 3)
	t.Logf("res: %d", res)
	res = activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5)
	t.Logf("res: %d", res)
	res = activityNotifications([]int32{1, 2, 3, 4, 4}, 4)
	t.Logf("res: %d", res)
}

func activityNotifications(expenditure []int32, d int32) int32 {
	var res int32
	freqTable := initFreqTable(expenditure, d)
	even := d%2 == 0

	target := d / 2
	if !even {
		target++
	}

	doubleMed := findMid(freqTable, target, even)
	fmt.Printf("double med:%d\n", doubleMed)

	for i := d; int(i) < len(expenditure); i++ {
		if doubleMed <= expenditure[i] {
			res++
		}
		freqTable[expenditure[int32(i)-d]]--
		freqTable[expenditure[i]]++
		doubleMed = findMid(freqTable, target, even)
		fmt.Printf("double med:%d\n", doubleMed)
	}

	return res
}

func findMid(freq map[int32]int32, target int32, isEven bool) int32 {
	var acc int32
	var med int32

	for i := 0; i <= 200; i++ {
		acc += freq[int32(i)]
		if acc >= target {
			med = int32(i)
			break
		}
	}

	if acc == target && isEven {
		var biggerMed int32
		for i := med + 1; i <= 200; i++ {
			if freq[i] >= 1 {
				biggerMed = i
				break
			}
		}
		return med + biggerMed
	}

	return med + med
}

func initFreqTable(expenditure []int32, d int32) map[int32]int32 {
	freq := make(map[int32]int32)

	for i := 0; i <= 200; i++ {
		freq[int32(i)] = 0
	}

	for _, num := range expenditure[:d] {
		freq[num]++
	}

	return freq
}
