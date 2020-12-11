package array

import (
	"sort"
	"strings"
	"testing"
)

func TestGroupAna(t *testing.T) {
	t.Log(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	mapStr := make(map[string][]string)
	res := [][]string{}

	for i := range strs {
		subStr := strings.Split(strs[i], "")
		sort.Strings(subStr)
		if cap(mapStr[strings.Join(subStr, "")]) == 0 {
			mapStr[strings.Join(subStr, "")] = []string{strs[i]}
		} else {
			mapStr[strings.Join(subStr, "")] = append(mapStr[strings.Join(subStr, "")], strs[i])
		}
	}

	for _, v := range mapStr {
		res = append(res, v)
	}
	return res
}
