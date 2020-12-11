package string

import "strconv"

func horspoolSearch(text string, pattern string, index int) (int, bool) {
	table := createTable(pattern)

	m := len(pattern) - 1
	found := false
	posR := index
	for posR < len(text) {
		k := 0
		for k <= m && text[posR-k] == pattern[m-k] {
			k++
		}
		if k == m+1 {
			found = true
			break
		}
		posR += table[string(text[posR])]
	}
	return posR, found
}

func createTable(pattern string) map[string]int {
	table := make(map[string]int)
	for i := 0; i <= 9; i++ {
		table[strconv.Itoa(i)] = len(pattern)
	}
	for i := 0; i < len(pattern)-1; i++ {
		table[string(pattern[i])] = len(pattern) - 1 - i
	}
	return table
}
