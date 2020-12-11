package dp

import (
	"testing"
)

func TestCoinCh1(t *testing.T) {
	t.Log(coinChange([]int{1, 2, 5}, 11))
	t.Log(coinChanges([]int{1, 2, 5}, 11))
}

// https://leetcode.com/problems/coin-change/

// dp[amount] = min(dp[amount], dp[amount-couns(i)]+1)
func coinChanges(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	coins = quicksort(coins)
	if coins[0] > amount {
		return -1
	}
	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = 1000000
	}
	dp[0] = 0

	// 	for a := 1; a <= amount; a++{
	// 		for _, coin := range coins{
	// 				if coin > a{
	// 						continue
	// 				}
	// 				dp[a] = min(dp[a-coin]+1, dp[a])
	// 		}
	// }

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			dp[a] = min(dp[a-coin]+1, dp[a])
		}
	}

	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	coins = quicksort(coins)
	if coins[0] > amount {
		return -1
	}

	dp := make([][]int, amount+1)
	for i := range dp {
		dp[i] = make([]int, len(coins))
	}
	dp[coins[0]][0] = 1
	if amount%coins[0] == 0 {
		dp[amount][0] = amount / coins[0]
	}

	if len(coins) > 1 {
		for a := coins[1]; a <= amount; a++ {
			for i, coin := range coins {
				if a-coin == 0 {
					dp[a][i] = 1
				} else {
					min := getMin(a-coin, dp, coins)
					if min > 0 {
						dp[a][i] = min + 1
					}
				}
			}
		}
	}

	ans := 100000
	for _, val := range dp[amount] {
		if val != 0 && val < ans {
			ans = val
		}
	}

	if ans == 100000 {
		return -1
	}

	return ans
}

func getMin(amount int, dp [][]int, coins []int) int {
	if amount < 0 {
		return 0
	}

	min := 10000000
	for i := range coins {
		if dp[amount][i] != 0 && dp[amount][i] < min {
			min = dp[amount][i]
		}
	}

	if min == 10000000 {
		return 0
	}
	return min
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
