package tester

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func DataToInt32(filename string) []int32 {
	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	flePath := filepath.Join(dir, "/data/"+filename)

	data, err := ioutil.ReadFile(flePath)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	temp := strings.Split(string(data), " ")
	res := make([]int32, len(temp))
	for i := range temp {
		x, _ := strconv.Atoi(temp[i])
		res[i] = int32(x)
	}
	return res
}

func DataToString(filename string) []string {
	data, _ := ioutil.ReadFile("../../" + filename)
	return strings.Split(string(data), " ")
}
