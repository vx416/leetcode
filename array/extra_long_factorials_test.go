package array

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestExtraLongFactorials(t *testing.T) {
	t.Logf("str: %+v", toSlice(54321))
	t.Logf("%+v", multiply([]byte{7, 2}, 222))
	extraLongFactorials(25)
}

func extraLongFactorials(n int32) {
	numSlice := toSlice(n)
	n = n - 1

	for n >= 1 {
		numSlice = multiply(numSlice, n)
		n--
	}
	print(numSlice)
}

func multiply(numSlice []byte, n int32) []byte {
	res := make([]byte, len(numSlice))

	var ext int32 = 0
	var d int32 = 0

	for i := range numSlice {
		d = (int32(numSlice[i])*n + ext) % 10
		ext = (int32(numSlice[i])*n + ext) / 10
		res[i] = byte(d)
	}

	for ext > 0 {
		d = ext % 10
		ext = ext / 10
		res = append(res, byte(d))
	}

	return res
}

func print(numSlice []byte) {
	var s strings.Builder

	n := len(numSlice) - 1
	for n >= 0 {
		c := strconv.Itoa(int(numSlice[n]))
		s.WriteString(c)
		n--
	}
	fmt.Println(s.String())
}

func toSlice(n int32) []byte {
	res := make([]byte, 0, 1)
	for int(n/10) != 0 {
		d := n % 10
		res = append(res, byte(d))
		n = n / 10
	}
	res = append(res, byte(n))

	return res
}
