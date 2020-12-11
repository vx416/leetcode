package sort

import (
	"fmt"
	"hackerrank/tester"
	"testing"
)

// source:https://www.hackerrank.com/challenges/lilys-homework/problem

func TestLilysHomework(t *testing.T) {
	t.Logf("res: %d", lilysHomework([]int32{7, 15, 12, 3}))
	t.Logf("res: %d", lilysHomework([]int32{3, 4, 2, 5, 1}))
	data := tester.DataToInt32("test.txt")
	t.Logf("resN:%d", lilysHomework(data))
	t.Logf("resN:%d", lilysHomework(tester.DataToInt32("test2.txt")))

}

func lilysHomework(arr []int32) int32 {
	sorted := make([]int32, len(arr))
	arr2 := make([]int32, len(arr))
	copy(sorted, arr)
	copy(arr2, arr)
	quicksort(sorted)

	index := make(map[int32]int)
	index2 := make(map[int32]int)
	for i := range arr {
		index[arr[i]] = i
		index2[arr[i]] = i
	}

	var swap int32 = 0
	var swap2 int32 = 0

	n := len(arr) - 1
	for i, val := range sorted {
		if index[val] != i {
			originVal := arr[i]
			valIndex := index[val]
			arr[i], arr[index[val]] = arr[index[val]], arr[i]
			index[originVal] = valIndex
			index[val] = i
			swap++
		}
		if index2[val] != n-i {
			originVal := arr2[n-i]
			valIndex := index2[val]
			arr2[n-i], arr2[index2[val]] = arr2[index2[val]], arr2[n-i]
			index2[originVal] = valIndex
			index2[val] = n - i
			swap2++
		}

		// if reflect.DeepEqual(arr, sorted) {
		// 	break
		// }
		// if sameReverse(arr2, sorted) {
		// 	break
		// }
	}

	if swap < swap2 {
		return swap
	}

	return swap2
}

func sameReverse(a []int32, b []int32) bool {
	n := len(b) - 1
	for i := range a {
		fmt.Printf("a:%d,b:%d\n", a[i], b[n-i])
		if a[i] != b[n-i] {
			return false
		}
	}
	return true
}
