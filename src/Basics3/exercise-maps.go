package main

import (
	"strings"
	"../../../../Go/misc/tour/src/golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, v := range(strings.Fields(s)) {
		ret[v] += 1
	}
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
