package dp

import (
	"fmt"
	"testing"
)

func TestCoinCh2(t *testing.T) {
	t.Log(change(5, []int{1, 2, 5}))
}

// https://leetcode.com/problems/coin-change-2/
// dp[amount, coin_j] = sum(i ~ 1~j dp[amount-coin_j, coin_i])
// dp[amonut] = sum(i ~ coins[i]<=amount, dp[amount-coins[i]])
// dp[4] = dp[3]+dp[2]
// dp[3] = dp[2]+dp[1]
// dp[amount, coin] += dp[amount-coin_i, coin_i] if amount > coin_i
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}
	quicksort(coins)
	if coins[0] > amount {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
		fmt.Printf("DP:%+v\n", dp)
	}

	// for a := 1; a <= amount; a++ {
	// 	for _, coin := range coins {
	// 		if coin <= a {
	// 			dp[a] += dp[a-coin]
	// 		}
	// 	}
	// }
	fmt.Printf("DP:%+v", dp)

	return dp[amount]

	// dp := make([][]int, amount+1)
	// for i := range dp {
	// 	dp[i] = make([]int, len(coins))
	// }

	// for a := coins[0]; a <= amount; a++ {
	// 	for i, coin := range coins {
	// 		if a-coin == 0 {
	// 			dp[a][i]++
	// 		} else if a-coin > 0 {
	// 			for j := 0; j <= i; j++ {
	// 				dp[a][i] += dp[a-coin][j]
	// 			}
	// 		}
	// 	}
	// }

	// sum := 0
	// for _, val := range dp[amount] {
	// 	sum += val
	// }
	// return sum
}
