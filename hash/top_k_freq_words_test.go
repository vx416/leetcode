package hash

import (
	"testing"
)

// https://leetcode.com/problems/top-k-frequent-words/

func TestTopKFreq(t *testing.T) {
	// t.Log(searchWord([]string{}, "a"))
	t.Log(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	t.Log(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}

func topKFrequent(words []string, k int) []string {
	wordFreq := make(map[string]int)
	for _, w := range words {
		wordFreq[w]++
	}

	max := 0
	freqWord := make(map[int][]string)
	for word, freq := range wordFreq {
		if freq > max {
			max = freq
		}
		if _, ok := freqWord[freq]; !ok {
			freqWord[freq] = make([]string, 0, 1)
		}
		i := searchWord(freqWord[freq], word)
		newStr := make([]string, 0, len(freqWord[freq])+1)
		newStr = append(newStr, freqWord[freq][:i]...)
		newStr = append(newStr, word)
		newStr = append(newStr, freqWord[freq][i:]...)
		freqWord[freq] = newStr
	}

	res := make([]string, 0, 5)

	n := max
	for len(res) < k {
		if ws, ok := freqWord[n]; ok {
			for _, w := range ws {
				res = append(res, w)
				if len(res) == k {
					break
				}
			}
		}
		n--
	}

	return res
}

func searchWord(words []string, word string) int {
	if len(words) == 0 {
		return 0
	}

	l := 0
	r := len(words) - 1

	for (r - l) > 1 {
		mid := l + ((r - l) / 2)
		t := words[mid]
		if t == word {
			return mid
		} else if word > t {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if words[r] < word {
		return r + 1
	}

	if words[l] < word {
		return l + 1
	}
	return l
}
