package array

import (
	"testing"
)

func TestClimbingTheLeaderBoard(t *testing.T) {
	scores := []int32{100, 100, 50, 40, 40, 20, 10}
	alices := []int32{5, 25, 50, 120}

	res := climbingLeaderboard(scores, alices)

	t.Logf("res: %+v", res)
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	currIndex := 0
	aliceTime := len(alice) - 1
	res := make([]int32, len(alice))

	var currRank int32 = 1

	for aliceTime >= 0 {
		for currIndex < len(scores) {
			if alice[aliceTime] < scores[currIndex] {
				currIndex++
				currRank++
				for currIndex < len(scores) && scores[currIndex-1] == scores[currIndex] {
					currIndex++
				}
			} else {
				res[aliceTime] = currRank
				break
			}
		}
		if currIndex == len(scores) {
			res[aliceTime] = currRank
		}

		aliceTime--
	}

	return res
}
