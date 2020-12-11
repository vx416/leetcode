package array

import (
	"strconv"
	"strings"
	"testing"
)

func TestInvalidTrans(t *testing.T) {
	t.Log(invalidTransactions([]string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"}))

	// t.Log(invalidTransactions([]string{"lee,886,1785,beijing", "alex,763,1157,amsterdam", "lee,277,129,amsterdam", "bob,770,105,amsterdam", "lee,603,926,amsterdam", "chalicefy,476,50,budapest", "lee,924,859,barcelona", "alex,302,590,amsterdam", "alex,397,1464,barcelona", "bob,412,1404,amsterdam", "lee,505,849,budapest"}))
}

// a>b>c
// c-a <= 60
// b-a <= 60

// https://leetcode.com/problems/invalid-transactions/
// ["alice,20,800,mtv","alice,50,100,beijing"]
func invalidTransactions(transactions []string) []string {
	ts := make([]trans, len(transactions))

	for i, t := range transactions {
		ts[i] = newTrans(t)
	}
	invalid := make([]string, 0, 1)
	set := addSet(&invalid)
	prevTrans := make(map[string][]trans)

	for i := range ts {
		if ts[i].amount > 1000 {
			set(ts[i].raw)
		}

		if len(prevTrans[ts[i].name]) == 0 {
			prevTrans[ts[i].name] = []trans{ts[i]}
		} else {
			for _, otherT := range prevTrans[ts[i].name] {
				if otherT.invalidWith(ts[i]) {
					set(otherT.raw)
					set(ts[i].raw)
				}
			}
			prevTrans[ts[i].name] = append(prevTrans[ts[i].name], ts[i])
		}
	}

	return invalid
}

func addSet(arr *[]string) func(s string) {
	exist := make(map[string]bool)

	return func(s string) {
		if _, ok := exist[s]; !ok {
			exist[s] = true
			*arr = append(*arr, s)
		}
	}
}

func newTrans(t string) trans {
	tokens := strings.Split(t, ",")
	name := tokens[0]
	min, _ := strconv.Atoi(tokens[1])
	amount, _ := strconv.Atoi(tokens[2])
	city := tokens[3]
	return trans{name, min, amount, city, t, false}
}

type trans struct {
	name    string
	min     int
	amount  int
	city    string
	raw     string
	invalid bool
}

func (t *trans) invalidWith(ot trans) bool {
	return abs(int64(ot.min-t.min)) <= 60 && t.city != ot.city
}
