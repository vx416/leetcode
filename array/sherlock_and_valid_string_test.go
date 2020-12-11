package array

import (
	"fmt"
	"testing"
)

// https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem

func TestSherlockAndValidString(t *testing.T) {
	t.Logf("ans: %s", isValid("aabbcd"))
	// t.Logf("ans: %s", isValid("aabbccddeefghi"))
	// t.Logf("ans: %s", isValid("abvc"))
	// t.Logf("ans: %s", isValid("abcc"))

	// t.Logf("ans: %s", isValid("abcdefghhgfedecba"))
}

func isValid(s string) string {

	charFreq := make(map[rune]int64)

	for _, char := range s {
		if _, ok := charFreq[char]; !ok {
			charFreq[char] = 0
		}
		charFreq[char]++
	}

	// var minFreq int64 = 10000
	// for _, freq := range charFreq {
	// 	if freq < minFreq {
	// 		minFreq = freq
	// 	}
	// }

	// var count int64
	// for _, freq := range charFreq {
	// 	if freq != minFreq {
	// 		count++
	// 	}
	// }

	// if count <= 1 {
	// 	return "YES"
	// }

	charCount := make(map[int64]int64)

	counts := make([]int64, 0, 1)
	for _, v := range charFreq {
		if _, ok := charCount[v]; !ok {
			charCount[v] = 0
			counts = append(counts, v)
		}
		charCount[v]++
	}

	fmt.Printf("key:%+v", counts)
	fmt.Printf("freqCount:%+v", charCount)

	if len(counts) == 1 {
		return "YES"
	}

	if len(counts) == 2 {
		var bigger, samller int64
		if counts[0] > counts[1] {
			bigger = counts[0]
			samller = counts[1]
		} else {
			bigger = counts[1]
			samller = counts[0]
		}
		if charCount[bigger] == 1 && abs(bigger-samller) == 1 {
			return "YES"
		}
		if charCount[samller] == 1 {
			return "YES"
		}

	}

	return "NO"
}

func abs(i int64) int64 {
	if i < 0 {
		return -1 * i
	}
	return i
}
