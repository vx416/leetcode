package dp

import (
	"testing"
)

func TestLenLIS(t *testing.T) {
	t.Log(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

// dp[i] = i-index 前所有的 increasing sequences
// dp[i] = for j in range i-1:
//            if try_join(dp[j], nums[i]): max_len = max(dp[j], max_len)
// https://leetcode.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([][2]int, len(nums))
	dp[0] = [2]int{1, nums[0]}
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		currMaxLen := 1
		for j := 0; j < i; j++ {
			if dp[j][1] < nums[i] && dp[j][0]+1 > currMaxLen {
				currMaxLen = dp[j][0] + 1
			}
		}
		dp[i] = [2]int{currMaxLen, nums[i]}
		if currMaxLen > maxLen {
			maxLen = currMaxLen
		}
	}

	return maxLen
}
