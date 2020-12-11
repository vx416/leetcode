package dp

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	t.Log(canJump([]int{2, 3, 1, 1, 4}))
	t.Log(canJump([]int{3, 2, 1, 0, 4}))
	t.Log(canJump([]int{3, 0, 8, 2, 0, 0, 1}))
}

// https://leetcode.com/problems/jump-game/
// dp[i] 代表 i 是 good index
// dp[i] = try(nums[i]) if dp[i-1]
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(dp); j++ {
			if dp[j] && nums[i] >= j-i {
				dp[i] = true
				break
			}
		}
		if !dp[i] && nums[i] >= len(nums)-1-i {
			dp[i] = true
		}
	}

	return dp[0]
}

func canJumpGreedy(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= len(nums)-1-i || nums[i] >= lastPos-i {
			lastPos = i
		}
	}
	return lastPos == 0
}
