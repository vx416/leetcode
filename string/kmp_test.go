package string

import (
	"testing"
)

func simpleMatch(s1, s2 string) string {
	var start int
	match := false
	for i, ch := range s1 {
		match = false
		for _, ch2 := range s2 {
			if ch != ch2 {
				break
			}
			match = true
		}
		if match {
			start = i
			break
		}
	}

	if match {
		return s1[start : start+len(s2)]
	}
	return ""
}

func TestMatch(t *testing.T) {
	t.Log(simpleMatch("abcd1234adsf", "cd123"))
	t.Log(KMPMatch("abbbccaac", "bbcca"))
	t.Log(KMPMatch("abcd1234adsf", "cd123"))
	t.Log(KMPMatch("asdfjzxcivjasvjidfaicvsdf", "jzxcivjasvji"))
}

func BenchmarkMatch(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			simpleMatch("sadfasdvjsdgjxklzcjvasdflcjxvasdfasdfjzxcivjasvjidfaicvsdf", "jzxcivjasvji")
		}
	})

	b.Run("kmp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KMPMatch("sadfasdvjsdgjxklzcjvasdflcjxvasdfasdfjzxcivjasvjidfaicvsdf", "jzxcivjasvji")
		}
	})

	b.Run("kmp2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KMPSearch("sadfasdvjsdgjxklzcjvasdflcjxvasdfasdfjzxcivjasvjidfaicvsdf", "jzxcivjasvji")
		}
	})

	b.Run("kmp2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			horspoolSearch("sadfasdvjsdgjxklzcjvasdflcjxvasdfasdfjzxcivjasvjidfaicvsdf", "jzxcivjasvji", 0)
		}
	})
}

// i :=> 目前比對到的 index of s1
// j :=> 起始比對的 index of s2
// next[i] :=> 代表 s2[i] 前面的字串 有n個字元是屬於 suffix == prefix
// when unmatch on i, j
// 	 從 next 找出 s2[j] 前面有幾個字元數用 suffix == prefix (next[j])
// 	 從 next 中我們可以得出 => s1[i-next[j]:i] == s2[0:next[j]]
//   因此在 s2 中我們只要從 next[j] 開始與 i 開始比對
// 	所以我們可以把 j 重新賦予值為 next[j]
func KMPMatch(s1, s2 string) string {
	next := getNext2(s2)

	i := 0
	j := 0
	match := false

	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			j = next[j]
			if j == 0 {
				i++
			}
		} else {
			j++
			i++
		}

		if j == len(s2) {
			match = true
			break
		}
	}

	if match {
		return s1[i-len(s2) : i]
	}

	return ""
}

func addI(i, j int, next []int) (int, int) {
	if next[j] == -1 {
		return i + 1, 0
	}
	return i + j - next[j] + j - 1, next[j]
}

func check(s1, s2 string, i, j int) (int, bool) {
	pre := j
	for {
		if i >= len(s1) {
			return j, false
		}
		if j >= len(s2) {
			return pre, true
		}

		if s1[i] != s2[j] {
			return j, false
		}
		i++
		j++
	}

}

func failFunc(s string) []int {
	failure := make([]int, len(s))
	failure[0] = -1
	i := 1
	j := 0
	for i < len(s)-1 {
		if j == -1 || s[j] == s[i] {
			i++
			j++
			failure[i] = j
		} else {
			j = failure[j]
		}
	}

	return failure
}

func getNext3(s string) []int {
	if len(s) == 0 {
		return []int{}
	}

	next := make([]int, len(s))
	next[0] = -1 // 代表前面沒有半個 char prefix == suffix

	prefixIndex, suffixIndex := -1, 0
	// 從 1 ~ n-1 個開始找 prefix == suffix
	for suffixIndex < len(s)-1 {

		if prefixIndex == -1 || s[suffixIndex] == s[prefixIndex] {
			// when s[suffixIndex] == s[prefixIndex]
			// next[suffixIndex+1] = prefixIndex+1
			// next[i] = 在 i index 前面有幾個char 符合 prefix == suffix
			prefixIndex++
			suffixIndex++
			next[suffixIndex] = prefixIndex
		} else {
			// when s[suffixIndex] != s[prefixIndex]
			// 代表 suffix 之前 有 prefix-1 個(k) char 一樣
			// 如果 k == 0 && s[suffixIndex] != s[prefixIndex]
			// 需要讓 suffixIndex+1 跟 0 開始重新比對
			k := prefixIndex
			if k == 0 {
				suffixIndex++
				prefixIndex = 0
				continue
			}

			// 如果 k > 0
			// 代表 suffixIndex 前 k 個字元跟 s[0:k] 一樣
			// 透過 next[k] 我們可以發現 k這個 Index 前有 next[k](m) 個字元 prefix == suffix
			// 同時也代表 suffixIndex 前 k 個字元也有 m 個字元 prefix == suffix
			// 因此我們可以透過將 prefixIndex 重新定位成 m 來過濾掉不必要的比對字元
			if k > 0 {
				m := next[k]
				prefixIndex = m
			}
		}
	}
	return next
}

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	// i 是當前的字元index位置
	// j 是 (i-1)index 中 有j個字元 prefix == suffix
	// j 是 要比對字元的起始位置

	i, j := 0, -1
	for i < len(s)-1 {
		// 如果 j == -1 代表 i == 0
		// 如果 s[i] == s[j]
		// s[1] == s[0]
		// next[2] = 1
		// s[2] == s[1]
		if j == -1 || s[i] == s[j] {
			i++
			j++
			// next[1] = 0
			next[i] = j
		} else {
			// 當 s[i] != s[j]
			// j = next[j]
			j = next[j]
		}
	}
	return next
}

func KMPSearch(S, T string) int {
	next := getNext(T)
	// fmt.Println(next)
	i := 0
	j := 0
	//同时满足才可以  找除字符串出现的第一个位置
	for i <= len(S)-1 && j <= len(T)-1 {

		if j == -1 || S[i] == T[j] {
			//当字符匹配时 i j 都加1
			i++
			j++
		} else {
			//子串的 偏移量 从next数组中取  i 不变
			j = next[j]
		}
	}
	//如果 j 大于 或者 等于 T串的长度 说明匹配成功
	if j >= len(T)-1 {
		return i - len(T) + 1
	}

	return 0
}

func getNext2(s string) []int {
	result := make([]int, len(s))
	result[0] = 0
	result[1] = 0
	if s[0] == s[1] {
		result[2] = 1
	}

	for i := 3; i < len(s); i++ {
		if result[i-1] != 0 {
			if result[i-1]+1 != i-1 && s[i-1] == s[result[i-1]] {
				result[i] = result[i-1] + 1
			} else {
				result[i] = 0
			}
		} else if result[i-1] == 0 {
			if s[i-1] == s[0] {
				result[i] = 1
			} else {
				result[i] = 0
			}
		}
	}

	return result
}
