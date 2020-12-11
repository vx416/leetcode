package dp

import (
	"testing"
)

func TestCanIwin(t *testing.T) {
	// t.Log(canIWin(10, 11))
	// t.Log(canIWin(4, 6))
	t.Log(canIWin(20, 210))
}

// dp[curr][chose]bool
func max2(chose []int, curr int, target int) bool {
	if chose[len(chose)-1]+curr >= target {
		return true
	}

	for i, val := range chose {
		newChose := make([]int, 0, len(chose))
		newChose = append(newChose, chose[i+1:]...)
		newChose = append(newChose, chose[:i]...)
		if !max2(newChose, curr+val, target) {
			return true
		}
	}

	return false

}

// dp[chooseSet, desiredTotaol]
// https://leetcode.com/problems/can-i-win/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	if maxChoosableInteger > desiredTotal {
		return true
	}

	dp := make(map[int]map[int]bool)
	res := win(maxChoosableInteger, desiredTotal, 0, dp)
	return res
}

func win(max, target, usedValues int, dp map[int]map[int]bool) bool {
	if target <= 0 {
		return false
	}

	if dp2, ok := dp[target]; ok {
		if res, ok := dp2[usedValues]; ok {
			return res
		}
	}

	for n := max; n >= 1; n-- {
		if (1 << n & usedValues) == 0 {
			if n >= target {
				return true
			}

			if !(win(max, target-n, usedValues|1<<n, dp)) {
				if _, ok := dp[target]; !ok {
					dp[target] = make(map[int]bool)
				}
				dp[target][usedValues] = true
				return true
			}
		}
	}
	if _, ok := dp[target]; !ok {
		dp[target] = make(map[int]bool)
	}
	dp[target][usedValues] = false
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
